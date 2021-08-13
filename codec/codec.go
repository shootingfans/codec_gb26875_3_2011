package codec

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"
	"github.com/shootingfans/codec_gb26875_3_2011/utils"

	"github.com/valyala/bytebufferpool"
)

var (
	ErrPacketInvalid         = errors.New("packet invalid")          // packet is invalid
	ErrPacketNotEnough       = errors.New("packet not enough")       // packet data not enough
	ErrPacketChecksumInvalid = errors.New("packet checksum invalid") // packet checksum invalid
)

const (
	HeadFlag          byte = 0x40 // packet header flag: 0x40 0x40
	TailFlag          byte = 0x23 // packet tail flag 0x23 0x23
	DefaultHeadLength int  = 26   // packet header length head_flag 2byte + serial 2byte + version 2byte + timestamp 6byte + source 6byte + target 6byte + data length 2byte
	DefaultTailLength int  = 3    // packet tail length check sum 1byte + tail flag 2byte
)

type Encoder interface {
	Encode(packet *constant.Packet) ([]byte, error)
}

type Decoder interface {
	Decode(b []byte) (*constant.Packet, int, error)
}

type Codec interface {
	Encoder
	Decoder
}

type myCodec struct{}

func (m myCodec) Encode(packet *constant.Packet) ([]byte, error) {
	buffer := bytebufferpool.Get()
	defer bytebufferpool.Put(buffer)
	// write head flag 0x40 0x40
	buffer.Write([]byte{HeadFlag, HeadFlag})
	// write serial id
	binary.Write(buffer, binary.LittleEndian, packet.Header.SerialId)
	// write version
	binary.Write(buffer, binary.BigEndian, packet.Header.Version)
	// write timestamp
	buffer.Write(utils.Timestamp2Bytes(packet.Header.Timestamp))
	// write source and target address
	address := make([]byte, 8)
	binary.LittleEndian.PutUint64(address, packet.Header.Source)
	buffer.Write(address[0:6])
	binary.LittleEndian.PutUint64(address, packet.Header.Target)
	buffer.Write(address[0:6])
	// write app data len
	binary.Write(buffer, binary.LittleEndian, uint16(len(packet.AppData)))
	// write action
	buffer.WriteByte(byte(packet.Action))
	// write app data
	buffer.Write(packet.AppData)
	// write checksum
	buffer.WriteByte(uint8(utils.Sum(buffer.Bytes()[2:])))
	// write tail flag 0x23 0x23
	buffer.Write([]byte{TailFlag, TailFlag})
	return buffer.Bytes(), nil
}

func (m myCodec) Decode(b []byte) (*constant.Packet, int, error) {
	if len(b) < DefaultHeadLength {
		return nil, 0, ErrPacketNotEnough
	}
	if b[0] != HeadFlag || b[1] != HeadFlag {
		return nil, 2, fmt.Errorf("%w : head flag %v invalid", ErrPacketInvalid, b[0:2])
	}
	// data length = action + app data
	dataLength := int(binary.LittleEndian.Uint16(b[24:26])) + 1
	packetLength := DefaultHeadLength + dataLength + DefaultTailLength
	if len(b) < packetLength {
		return nil, 0, ErrPacketNotEnough
	}
	if b[packetLength-2] != TailFlag || b[packetLength-1] != TailFlag {
		return nil, packetLength, fmt.Errorf("%w : tail flag %v invalid", ErrPacketInvalid, b[packetLength-2:packetLength])
	}
	if sum := uint8(utils.Sum(b[2 : packetLength-DefaultTailLength])); sum != b[packetLength-DefaultTailLength] {
		return nil, packetLength, fmt.Errorf("%w : checksum fail %v != %v", ErrPacketChecksumInvalid, sum, b[packetLength-DefaultTailLength])
	}
	packet := constant.Packet{
		Header: constant.Header{
			SerialId:  binary.LittleEndian.Uint16(b[2:4]),
			Version:   constant.Version(binary.BigEndian.Uint16(b[4:6])),
			Timestamp: utils.Bytes2Timestamp(b[6:12]),
		},
		Action:  constant.Action(b[26]),
		AppData: b[27 : packetLength-DefaultTailLength],
	}
	address := make([]byte, 8)
	copy(address[0:6], b[12:18])
	packet.Header.Source = binary.LittleEndian.Uint64(address)
	copy(address[0:6], b[18:24])
	packet.Header.Target = binary.LittleEndian.Uint64(address)
	return &packet, packetLength, nil
}
