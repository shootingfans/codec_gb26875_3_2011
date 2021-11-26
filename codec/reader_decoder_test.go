package codec

import (
	"bytes"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"
	"github.com/shootingfans/codec_gb26875_3_2011/utils"

	"github.com/stretchr/testify/assert"
)

func TestReaderDecoder(t *testing.T) {
	t.Run("test close when C not empty", func(t *testing.T) {
		buf := bytes.NewBuffer(make([]byte, 0))
		rd := NewReaderDecoder(buf)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			for range [2][]int{} {
				select {
				case p, ok := <-rd.C:
					if !ok {
						return
					}
					assert.EqualValues(t, p.Header, constant.Header{
						Version:   0x101,
						Timestamp: utils.Bytes2Timestamp([]byte{0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14}),
						Target:    0x010203040506,
					})
					assert.Equal(t, p.Action, constant.AckAction)
				}
			}
		}()
		go func() {
			defer wg.Done()
			bb := []byte{
				0x00, 0x00, 0x01, 0x01, 0x18, 0x0d, 0x11, 0x16, 0x0a, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00, 0x00, 0x03,
			}
			buf.Write([]byte{0x40, 0x40})
			buf.Write(bb)
			buf.Write([]byte{0x11, 0x23, 0x24})
			bb = append(bb, uint8(utils.Sum(bb)), 0x23, 0x23)
			for range [3][]int{} {
				buf.Write([]byte{0x40, 0x40})
				buf.Write(bb)
				time.Sleep(time.Millisecond * 10)
			}
		}()
		time.Sleep(time.Millisecond * 500)
		rd = nil
		runtime.GC()
		time.Sleep(time.Millisecond * 100)
		wg.Wait()
	})
	t.Run("test close", func(t *testing.T) {
		buf := bytes.NewBuffer(make([]byte, 0))
		buf.Write([]byte{0x40, 0x40, 0x01, 0x00, 0x00, 0x01, 0x05, 0x00, 0x09, 0x02, 0x07, 0x15, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00, 0x0b, 0x0a, 0x09, 0x08, 0x07, 0x06, 0x03, 0x00, 0x04, 0x59, 0x01, 0x00, 0xd1, 0x23, 0x23})
		rd := NewReaderDecoder(buf)
		time.Sleep(time.Millisecond * 300)
		time.Sleep(time.Second)
		rd.Close()
	})
}
