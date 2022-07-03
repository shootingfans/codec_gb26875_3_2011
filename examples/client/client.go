package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/codec"
	"github.com/shootingfans/codec_gb26875_3_2011/utils"
)

var serverAddr = flag.String("server-host", "127.0.0.1:8181", "server address")
var serialNumber uint16 = 0

func main() {
	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connected to ", *serverAddr)
	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGALRM)
	defer conn.Close()
	ticker := time.NewTicker(time.Second * 1)
	buf := bytes.NewBuffer(make([]byte, 100))
	cd := codec.NewReaderDecoder(conn)
	defer cd.Close()
	for {
		select {
		case s := <-sg:
			log.Println("stopped on:", s)
			ticker.Stop()
			return
		case <-ticker.C:
			buf.Reset()
			buf.Write([]byte{0x40, 0x40})
			serialNumber++
			binary.Write(buf, binary.BigEndian, serialNumber)
			buf.Write([]byte{0x01, 0x01})
			buf.Write(utils.Timestamp2Bytes(time.Now().Unix()))
			buf.Write(bytes.Repeat([]byte{0x00}, 12))
			buf.Write([]byte{0x08, 0x00, 0x02, 0x1c, 0x01})
			buf.Write(utils.Timestamp2Bytes(time.Now().Unix()))
			buf.Write([]byte{uint8(utils.Sum(buf.Bytes()[2:35])), 0x23, 0x23})
			conn.SetWriteDeadline(time.Now().Add(time.Second))
			by := buf.Bytes()
			n, err := buf.WriteTo(conn)
			if err != nil {
				if err == io.EOF {
					return
				}
				log.Printf("write to server fail: %v\n", err)
			}
			if n > 0 {
				log.Printf("write %d bytes to server %#x", n, by[0:n])
			}
		case p, ok := <-cd.C:
			if !ok {
				return
			}
			log.Printf("receive packet is empty?: %v\n", p.IsEmpty())
			log.Printf("recevie packet from server: action = %s\n", p.Action.Name())
		}
	}
}
