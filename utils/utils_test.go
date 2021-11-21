package utils_test

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/utils"

	"github.com/stretchr/testify/assert"
)

func TestB2S(t *testing.T) {
	b := []byte{0x61, 0x73, 0x64, 0x66, 0x31, 0x32, 0x33, 0x34, 0x35}
	assert.Equal(t, utils.B2S(b), "asdf12345")
}

func TestS2B(t *testing.T) {
	s := "123456"
	assert.Equal(t, utils.S2B(s), []byte(s))
}

func TestEncodeGB18030(t *testing.T) {
	b := []byte("这是一个转码测试")
	assert.EqualValues(t, []byte{0xD5, 0xE2, 0xCA, 0xC7, 0xD2, 0xBB, 0xB8, 0xF6, 0xD7, 0xAA, 0xC2, 0xEB, 0xB2, 0xE2, 0xCA, 0xD4}, utils.EncodeGB18030(b))
}

func TestDecodeGB18030(t *testing.T) {
	b := []byte{0xd5, 0xe2, 0xca, 0xc7, 0xd2, 0xbb, 0xb8, 0xf6, 0xd7, 0xaa, 0xc2, 0xeb, 0xb2, 0xe2, 0xca, 0xd4}
	assert.EqualValues(t, []byte("这是一个转码测试"), utils.DecodeGB18030(b))
}

func TestBytes2Int16(t *testing.T) {
	t.Run("test big order", func(t *testing.T) {
		assert.Equal(t, utils.Bytes2Int16([]byte{0xfe, 0x11}, binary.BigEndian), int16(-495))
	})
	t.Run("test little order", func(t *testing.T) {
		assert.Equal(t, utils.Bytes2Int16([]byte{0xfe, 0x11}, binary.LittleEndian), int16(4606))
	})
}

func TestSum(t *testing.T) {
	res := utils.Sum([]byte{0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14, 0x30, 0x00, 0x02, 0x02, 0x01, 0x03, 0x00, 0xd9, 0x00, 0x06, 0x00, 0x02, 0x00, 0xa3, 0xc1, 0xc7, 0xf8, 0xa3, 0xb1, 0xb2, 0xe3, 0xdf, 0xc8, 0xb2, 0xb8, 0xd7, 0xdf, 0xc0, 0xc8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x12, 0x13, 0x01, 0x08, 0x14})
	t.Run("test one sum", func(t *testing.T) {
		assert.True(t, res.Equal([]byte{0x50}, binary.BigEndian))
	})
	t.Run("test tow sum", func(t *testing.T) {
		res.Equal([]byte{0x0e, 0x50}, binary.BigEndian)
		assert.True(t, res.Equal([]byte{0x50, 0x0e}, binary.LittleEndian))
	})
	t.Run("test other length", func(t *testing.T) {
		assert.False(t, res.Equal([]byte{}, binary.BigEndian))
	})
}

func TestTimestamp2Bytes(t *testing.T) {
	tm1 := time.Date(2021, 1, 3, 15, 30, 41, 0, time.Local)
	assert.Equal(t, utils.Timestamp2Bytes(tm1.Unix()), []byte{41, 30, 15, 3, 1, 21})
}

func TestBytes2Timestamp(t *testing.T) {
	by := []byte{0, 0, 15, 3, 1, 21}
	tm := time.Date(2021, 1, 3, 15, 0, 0, 0, time.Local)
	assert.Equal(t, utils.Bytes2Timestamp(by), tm.Unix())
	assert.Equal(t, utils.Bytes2Timestamp(by[2:]), tm.Unix())
}

func BenchmarkB2S(b *testing.B) {
	by := []byte("12345")
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		utils.B2S(by)
	}
}

func BenchmarkS2B(b *testing.B) {
	s := "12345"
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		utils.S2B(s)
	}
}
