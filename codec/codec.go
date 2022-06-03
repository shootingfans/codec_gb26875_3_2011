// Package codec is define Encoder and Decoder for GB26875.3-2011
package codec

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"
	"github.com/shootingfans/codec_gb26875_3_2011/utils"

	"github.com/valyala/bytebufferpool"
)

// define errors
var (
	ErrPacketInvalid         = errors.New("packet invalid")          // packet is invalid
	ErrPacketNotEnough       = errors.New("packet not enough")       // packet data not enough
	ErrPacketChecksumInvalid = errors.New("packet checksum invalid") // packet checksum invalid

)

// define default codec
var defaultCodec = myCodec{}

const (
	HeadFlag          byte = 0x40 // packet header flag: 0x40 0x40
	TailFlag          byte = 0x23 // packet tail flag 0x23 0x23
	DefaultHeadLength int  = 26   // packet header length head_flag 2byte + serial 2byte + version 2byte + timestamp 6byte + source 6byte + target 6byte + data length 2byte
	DefaultTailLength int  = 3    // packet tail length check sum 1byte + tail flag 2byte
)

// Encode is encoded GB26875.3-2011 packet to bytes
func Encode(packet *constant.Packet) ([]byte, error) {
	return defaultCodec.Encode(packet)
}

// Decode is decoded bytes to GB26875.3-2011 packet
func Decode(b []byte) (*constant.Packet, int, error) {
	return defaultCodec.Decode(b)
}

// Encoder is encoded packet to bytes
type Encoder interface {
	Encode(packet *constant.Packet) ([]byte, error)
}

// Decoder is decoded bytes to packet GB26875.3-2011
type Decoder interface {
	Decode(b []byte) (*constant.Packet, int, error)
}

// Codec is Encoder and Decoder of GB26875.3-2011
type Codec interface {
	Encoder
	Decoder
}

// AppDecoder is app data decoder
type AppDecoder interface {
	Decode(b []byte, packet *constant.Packet)
}

// AppDataDecoder is decode app data func
type AppDataDecoder func(b []byte, packet *constant.Packet)

// Decode is decode app data
func (a AppDataDecoder) Decode(b []byte, packet *constant.Packet) {
	a(b, packet)
}

var (
	// typeDecoders define decoder of all app type
	typeDecoders = map[constant.AppType]AppDecoder{
		constant.UploadSystemStateAppType:                      AppDataDecoder(decodeUploadSystemState),
		constant.UploadEquipmentStateAppType:                   AppDataDecoder(decodeUploadEquipmentState),
		constant.UploadEquipmentParameterAppType:               AppDataDecoder(decodeUploadEquipmentParameter),
		constant.UploadSystemOperatingInformationAppType:       AppDataDecoder(decodeUploadSystemOperatingInformation),
		constant.UploadSystemSoftwareVersionAppType:            AppDataDecoder(decodeUploadSystemSoftwareVersion),
		constant.UploadSystemConfigureAppType:                  AppDataDecoder(decodeUploadSystemConfigure),
		constant.UploadEquipmentConfigureAppType:               AppDataDecoder(decodeUploadEquipmentConfigure),
		constant.UploadSystemTimeAppType:                       AppDataDecoder(decodeUploadSystemTime),
		constant.UploadTransmissionStateAppType:                AppDataDecoder(decodeUploadTransmissionState),
		constant.UploadTransmissionOperatingInformationAppType: AppDataDecoder(decodeUploadTransmissionOperatingInformation),
		constant.UploadTransmissionSoftwareVersionAppType:      AppDataDecoder(decodeUploadTransmissionSoftwareVersion),
		constant.UploadTransmissionConfigureAppType:            AppDataDecoder(decodeUploadTransmissionConfigure),
		constant.UploadTransmissionTimeAppType:                 AppDataDecoder(decodeUploadTransmissionTime),
	}
)

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
		return nil, 1, fmt.Errorf("%w : head flag %v invalid", ErrPacketInvalid, b[0:2])
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
		Action: constant.Action(b[26]),
	}
	if dataLength > 0 {
		packet.AppData = b[27 : packetLength-DefaultTailLength]
	}
	address := make([]byte, 8)
	copy(address[0:6], b[12:18])
	packet.Header.Source = binary.LittleEndian.Uint64(address)
	copy(address[0:6], b[18:24])
	packet.Header.Target = binary.LittleEndian.Uint64(address)
	switch packet.Action {
	case constant.SendDataAction, constant.ResponseAction:
		DecodeAppData(&packet)
	}
	return &packet, packetLength, nil
}

// DecodeAppData decode the application data
func DecodeAppData(packet *constant.Packet) {
	if len(packet.AppData) < 2 {
		return
	}
	appType, count := constant.AppType(packet.AppData[0]), int(packet.AppData[1])
	decoder, ok := typeDecoders[appType]
	if !ok {
		return
	}
	perLength := (len(packet.AppData) - 2) / count
	for i := 0; i < count; i++ {
		decoder.Decode(packet.AppData[2+i*perLength:2+i*perLength+perLength], packet)
	}
	return
}

// RegistryAppDecoder registry app decoder
func RegistryAppDecoder(tp constant.AppType, fn AppDecoder) (overwrite bool) {
	_, overwrite = typeDecoders[tp]
	typeDecoders[tp] = fn
	return
}

// decodeUploadSystemState decode upload system state
func decodeUploadSystemState(b []byte, packet *constant.Packet) {
	if len(b) < 10 {
		return
	}
	packet.ControllerStates = append(packet.ControllerStates, constant.ControllerStateInfo{
		Ctrl:      constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
		Flag:      constant.StateFlag(binary.LittleEndian.Uint16(b[2:4])),
		Timestamp: utils.Bytes2Timestamp(b[4:10]),
	})
}

// decodeUploadEquipmentState decode upload equipment state
func decodeUploadEquipmentState(b []byte, packet *constant.Packet) {
	if len(b) < 46 {
		return
	}
	packet.EquipmentStates = append(packet.EquipmentStates, constant.EquipmentStateInfo{
		Equ: constant.Equipment{
			Ctrl: constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
			Type: constant.EquipmentType(b[2]),
			Addr: constant.EquipmentAddr(binary.LittleEndian.Uint32(b[3:7])),
		},
		Flag:        constant.StateFlag(binary.LittleEndian.Uint16(b[7:9])),
		Description: utils.B2S(utils.DecodeGB18030(b[9:40])),
		Timestamp:   utils.Bytes2Timestamp(b[40:46]),
	})
}

// decodeUploadEquipmentParameter decode upload equipment parameter
func decodeUploadEquipmentParameter(b []byte, packet *constant.Packet) {
	if len(b) < 16 {
		return
	}
	packet.EquipmentParameters = append(packet.EquipmentParameters, constant.EquipmentParameterInfo{
		Equ: constant.Equipment{
			Ctrl: constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
			Type: constant.EquipmentType(b[2]),
			Addr: constant.EquipmentAddr(binary.LittleEndian.Uint32(b[3:7])),
		},
		Info:      constant.ParameterInfo{Type: constant.ParameterType(b[7]), Value: constant.NewParameterValue(b[8:10])},
		Timestamp: utils.Bytes2Timestamp(b[10:16]),
	})
}

// decodeUploadSystemOperatingInformation decode upload system operating information
func decodeUploadSystemOperatingInformation(b []byte, packet *constant.Packet) {
	if len(b) < 10 {
		return
	}
	packet.ControllerOperations = append(packet.ControllerOperations, constant.ControllerOperationInfo{
		Ctrl:      constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
		Flag:      constant.OperationFlag(b[2]),
		Operator:  int(b[3]),
		Timestamp: utils.Bytes2Timestamp(b[4:10]),
	})
}

// decodeUploadSystemSoftwareVersion decode upload software version
func decodeUploadSystemSoftwareVersion(b []byte, packet *constant.Packet) {
	if len(b) < 4 {
		return
	}
	packet.ControllerVersions = append(packet.ControllerVersions, constant.ControllerVersion{
		Ctrl:    constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
		Version: constant.Version(binary.BigEndian.Uint16(b[2:4])),
	})
}

// decodeUploadSystemConfigure decode upload system configure
func decodeUploadSystemConfigure(b []byte, packet *constant.Packet) {
	if len(b) < 3 {
		return
	}
	length := int(b[2])
	if len(b) < 3+length {
		return
	}
	packet.ControllerConfigures = append(packet.ControllerConfigures, constant.ControllerConfigure{
		Ctrl:      constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
		Configure: utils.B2S(utils.DecodeGB18030(b[3 : 3+length])),
	})
}

// decodeUploadEquipmentConfigure decode upload equipment configure
func decodeUploadEquipmentConfigure(b []byte, packet *constant.Packet) {
	if len(b) < 38 {
		return
	}
	packet.EquipmentConfigures = append(packet.EquipmentConfigures, constant.EquipmentConfigure{
		Equ: constant.Equipment{
			Ctrl: constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
			Type: constant.EquipmentType(b[2]),
			Addr: constant.EquipmentAddr(binary.LittleEndian.Uint32(b[3:7])),
		},
		Description: utils.B2S(utils.DecodeGB18030(b[7:38])),
	})
}

// decodeUploadSystemTime decode upload system time
func decodeUploadSystemTime(b []byte, packet *constant.Packet) {
	if len(b) < 8 {
		return
	}
	packet.ControllerTimestamps = append(packet.ControllerTimestamps, constant.ControllerTimestamp{
		Ctrl:      constant.Controller{Type: constant.ControllerType(b[0]), Addr: int(b[1])},
		Timestamp: utils.Bytes2Timestamp(b[2:8]),
	})
}

// decodeUploadTransmissionState decode upload transmission state
func decodeUploadTransmissionState(b []byte, packet *constant.Packet) {
	if len(b) < 7 {
		return
	}
	packet.TransmissionStates = append(packet.TransmissionStates, constant.TransmissionStateInfo{
		Flag:      constant.StateFlag(b[0]),
		Timestamp: utils.Bytes2Timestamp(b[1:7]),
	})
}

// decodeUploadTransmissionOperatingInformation decode upload transmission operating information
func decodeUploadTransmissionOperatingInformation(b []byte, packet *constant.Packet) {
	if len(b) < 8 {
		return
	}
	packet.TransmissionOperations = append(packet.TransmissionOperations, constant.TransmissionOperationInfo{
		Flag:      constant.OperationFlag(b[0]),
		Operator:  int(b[1]),
		Timestamp: utils.Bytes2Timestamp(b[2:8]),
	})
}

// decodeUploadTransmissionSoftwareVersion upload transmission software version
func decodeUploadTransmissionSoftwareVersion(b []byte, packet *constant.Packet) {
	if len(b) < 2 {
		return
	}
	packet.TransmissionVersions = append(packet.TransmissionVersions, constant.TransmissionVersion{Version: constant.Version(binary.BigEndian.Uint16(b[0:2]))})
}

// decodeUploadTransmissionConfigure decode upload transmission configure
func decodeUploadTransmissionConfigure(b []byte, packet *constant.Packet) {
	if len(b) < 1 {
		return
	}
	length := int(b[0])
	if len(b) < length+1 {
		return
	}
	packet.TransmissionConfigures = append(packet.TransmissionConfigures, constant.TransmissionConfigure{Configure: utils.B2S(utils.DecodeGB18030(b[1 : 1+length]))})
}

// decodeUploadTransmissionTime decode upload transmission time
func decodeUploadTransmissionTime(b []byte, packet *constant.Packet) {
	if len(b) < 6 {
		return
	}
	packet.TransmissionTimestamps = append(packet.TransmissionTimestamps, constant.TransmissionTimestamp{Timestamp: utils.Bytes2Timestamp(b[0:6])})
}

// NewQuerySystemStateAppData create a query system state app data request
func NewQuerySystemStateAppData(controllers ...constant.Controller) []byte {
	if len(controllers) == 0 {
		return nil
	}
	b := make([]byte, 2+len(controllers)*2)
	b[0], b[1] = byte(constant.QuerySystemStateAppType), byte(len(controllers))
	for idx, controller := range controllers {
		offset := idx * 2
		b[2+offset], b[3+offset] = byte(controller.Type), byte(controller.Addr)
	}
	return b
}

// NewQueryEquipmentStateAppData create a query equipment state app data request
func NewQueryEquipmentStateAppData(equipments ...constant.Equipment) []byte {
	if len(equipments) == 0 {
		return nil
	}
	b := make([]byte, 2+len(equipments)*6)
	b[0], b[1] = byte(constant.QueryEquipmentStateAppType), byte(len(equipments))
	for idx, equipment := range equipments {
		offset := idx * 6
		b[2+offset], b[3+offset] = byte(equipment.Ctrl.Type), byte(equipment.Ctrl.Addr)
		binary.LittleEndian.PutUint32(b[4+offset:8+offset], uint32(equipment.Addr))
	}
	return b
}

// NewQueryEquipmentParameterAppData create a query equipment parameter app data request
func NewQueryEquipmentParameterAppData(equipments ...constant.Equipment) []byte {
	if len(equipments) == 0 {
		return nil
	}
	b := make([]byte, 2+len(equipments)*6)
	b[0], b[1] = byte(constant.QueryEquipmentParameterAppType), byte(len(equipments))
	for idx, equipment := range equipments {
		offset := idx * 6
		b[2+offset], b[3+offset] = byte(equipment.Ctrl.Type), byte(equipment.Ctrl.Addr)
		binary.LittleEndian.PutUint32(b[4+offset:8+offset], uint32(equipment.Addr))
	}
	return b
}

// NewQuerySystemOperatingInformationAppData create a query system operating information app data request
func NewQuerySystemOperatingInformationAppData(controller constant.Controller, total int, startTime time.Time) []byte {
	return append([]byte{byte(constant.QuerySystemOperatingInformationAppType), 0x01, byte(controller.Type), byte(controller.Addr), byte(total)}, utils.Timestamp2Bytes(startTime.Unix())...)
}

// NewQuerySystemSoftwareVersionAppData create a query system software version app data request
func NewQuerySystemSoftwareVersionAppData(controller constant.Controller) []byte {
	return []byte{byte(constant.QuerySystemSoftwareVersionAppType), 0x01, byte(controller.Type), byte(controller.Addr)}
}

// NewQuerySystemConfigureAppData create a query system configure app data request
func NewQuerySystemConfigureAppData(controllers ...constant.Controller) []byte {
	if len(controllers) == 0 {
		return nil
	}
	b := make([]byte, 2+len(controllers)*2)
	b[0], b[1] = byte(constant.QuerySystemConfigureAppType), byte(len(controllers))
	for idx, controller := range controllers {
		offset := idx * 2
		b[2+offset], b[3+offset] = byte(controller.Type), byte(controller.Addr)
	}
	return b
}

// NewQueryEquipmentConfigureAppData create a query equipment configure app data request
func NewQueryEquipmentConfigureAppData(equipments ...constant.Equipment) []byte {
	if len(equipments) == 0 {
		return nil
	}
	b := make([]byte, 2+len(equipments)*6)
	b[0], b[1] = byte(constant.QueryEquipmentConfigureAppType), byte(len(equipments))
	for idx, equipment := range equipments {
		offset := idx * 6
		b[2+offset], b[3+offset] = byte(equipment.Ctrl.Type), byte(equipment.Ctrl.Addr)
		binary.LittleEndian.PutUint32(b[4+offset:8+offset], uint32(equipment.Addr))
	}
	return b
}

// NewQuerySystemTimeAppData create a query system time app data request
func NewQuerySystemTimeAppData(controller constant.Controller) []byte {
	return []byte{byte(constant.QuerySystemTimeAppType), 1, byte(controller.Type), byte(controller.Addr)}
}

// NewQueryTransmissionStateAppData create a query transmission state app data request
func NewQueryTransmissionStateAppData() []byte {
	return []byte{byte(constant.QueryTransmissionStateAppType), 0x01, 0x00}
}

// NewQueryTransmissionOperatingInformationAppData create a query transmission operating information app data request
func NewQueryTransmissionOperatingInformationAppData(total int, startTime time.Time) []byte {
	return append([]byte{byte(constant.QueryTransmissionOperatingInformationAppType), 0x01, byte(total)}, utils.Timestamp2Bytes(startTime.Unix())...)
}

// NewQueryTransmissionSoftwareVersionAppData create a query transmission software version app data request
func NewQueryTransmissionSoftwareVersionAppData() []byte {
	return []byte{byte(constant.QueryTransmissionSoftwareVersionAppType), 0x01, 0x00}
}

// NewQueryTransmissionConfigureAppData create a query transmission configure app data request
func NewQueryTransmissionConfigureAppData() []byte {
	return []byte{byte(constant.QueryTransmissionConfigureAppType), 0x01, 0x00}
}

// NewQueryTransmissionTimeAppData create a query transmission time app data request
func NewQueryTransmissionTimeAppData() []byte {
	return []byte{byte(constant.QueryTransmissionTimeAppType), 0x01, 0x00}
}

// NewInitializeTransmissionAppData create a initialize transmission app data request
func NewInitializeTransmissionAppData() []byte {
	return []byte{byte(constant.InitializeTransmissionAppType), 0x01, 0x00}
}

// NewSyncTransmissionTimeAppData create a sync transmission app data request
func NewSyncTransmissionTimeAppData(syncTime time.Time) []byte {
	return append([]byte{byte(constant.SyncTransmissionTimeAppType), 0x01}, utils.Timestamp2Bytes(syncTime.Unix())...)
}

// NewInspectSentriesAppData create a inspect sentries app data request
func NewInspectSentriesAppData(timeoutMinute int) []byte {
	return []byte{byte(constant.InspectSentriesAppType), 0x01, byte(timeoutMinute)}
}
