// Package utils is define utils function
package utils

import (
	"bytes"
	"encoding/binary"
	"time"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// B2S use no copy convert byte slice to string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B use no copy convert string to byte slice
func S2B(s string) []byte {
	p := (*[2]uintptr)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&[3]uintptr{p[0], p[1], p[1]}))
}

// DecodeGB18030 is decode GB18030
func DecodeGB18030(b []byte) []byte {
	eDesc := make([]byte, 300)
	cnt, _, _ := simplifiedchinese.GB18030.NewDecoder().Transform(eDesc, b, false)
	return bytes.TrimRight(bytes.TrimSpace(eDesc[0:cnt]), "\x00")
}

// EncodeGB18030 is encode GB18030
func EncodeGB18030(b []byte) []byte {
	dst := make([]byte, 300)
	cnt, _, _ := simplifiedchinese.GB18030.NewEncoder().Transform(dst, b, false)
	return dst[0:cnt]
}

// Bytes2Int16 is convert byte slice to int16
func Bytes2Int16(b []byte, order binary.ByteOrder) (value int16) {
	binary.Read(bytes.NewReader(b), order, &value)
	return
}

// Sum is get byte slice sum
func Sum(b []byte) SumResult {
	var a uint32
	for _, c := range b {
		a += uint32(c)
	}
	return SumResult(a)
}

// SumResult is sum result
type SumResult uint32

// Bytes16 is get sum result convert to uint16 byte slice
func (s SumResult) Bytes16(order binary.ByteOrder) []byte {
	b := make([]byte, 2)
	if order == binary.LittleEndian {
		binary.LittleEndian.PutUint16(b, uint16(s))
	} else {
		binary.BigEndian.PutUint16(b, uint16(s))
	}
	return b
}

// Equal is compare byte slice sum
func (s SumResult) Equal(want []byte, order binary.ByteOrder) bool {
	switch len(want) {
	case 1:
		return uint8(s) == want[0]
	case 2:
		m := s.Bytes16(order)
		return m[0] == want[0] && m[1] == want[1]
	default:
		return false
	}
}

// Bytes2Timestamp is convert bytes to second timestamp
func Bytes2Timestamp(b []byte) int64 {
	if len(b) < 6 {
		b = append(make([]byte, 6-len(b)), b...)
	}
	return time.Date(2000+int(b[5]), time.Month(int(b[4])), int(b[3]), int(b[2]), int(b[1]), int(b[0]), 0, time.Local).Unix()
}

// Timestamp2Bytes is convert second timestamp to bytes
func Timestamp2Bytes(timestamp int64) []byte {
	tm := time.Unix(timestamp, 0)
	return []byte{uint8(tm.Second()), uint8(tm.Minute()), uint8(tm.Hour()), uint8(tm.Day()), uint8(tm.Month()), uint8(tm.Year() - 2000)}
}
