package codec

import (
	"bufio"
	"errors"
	"io"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"
)

// ReaderDecoder is a decoder for io.Reader
type ReaderDecoder struct {
	*readerDecoder
}

type readerDecoder struct {
	reader       *bufio.Reader
	close        chan struct{}
	closed       uint32
	C            chan *constant.Packet
	readDuration time.Duration
}

func (r *readerDecoder) Close() error {
	if atomic.CompareAndSwapUint32(&r.closed, 0, 1) {
		close(r.close)
	}
	return nil
}

func (r *readerDecoder) cron() {
	defer close(r.C)
	for {
		select {
		case <-r.close:
			return
		default:
			b, _ := r.reader.Peek(256)
			p, n, err := Decode(b)
			if err != nil {
				if errors.Is(err, ErrPacketInvalid) || errors.Is(err, ErrPacketChecksumInvalid) {
					r.reader.Discard(n)
				}
				continue
			}
			select {
			case <-r.close:
				return
			case r.C <- p:
			}
		}
	}
}

// NewReaderDecoder is create ReaderDecoder
func NewReaderDecoder(reader io.Reader) *ReaderDecoder {
	rc := &ReaderDecoder{
		&readerDecoder{
			reader: bufio.NewReader(reader),
			close:  make(chan struct{}),
			C:      make(chan *constant.Packet),
		},
	}
	go rc.cron()
	runtime.SetFinalizer(rc, func(r *ReaderDecoder) {
		_ = r.Close()
	})
	return rc
}
