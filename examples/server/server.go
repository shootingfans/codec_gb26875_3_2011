package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/shootingfans/codec_gb26875_3_2011/codec"
	"github.com/shootingfans/codec_gb26875_3_2011/constant"
)

var listenOn = flag.String("listen", "127.0.0.1:8181", "listen address")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", *listenOn)
	if err != nil {
		log.Fatalln(err)
	}
	var wg sync.WaitGroup
	var connections sync.Map
	wg.Add(1)
	go func() {
		defer wg.Done()
		sg := make(chan os.Signal)
		signal.Notify(sg, syscall.SIGINT, syscall.SIGALRM)
		log.Println("Stopped: ", <-sg)
		listener.Close()
	}()
	log.Println("listen on :", *listenOn)
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		wg.Add(1)
		go processConnection(&connections, conn, &wg)
	}
	connections.Range(func(_, value interface{}) bool {
		value.(net.Conn).Close()
		return true
	})
}

func processConnection(storage *sync.Map, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
	storage.Store(conn.RemoteAddr().Network(), conn)
	defer storage.Delete(conn)
	cd := codec.NewReaderDecoder(conn)
	defer cd.Close()
	for {
		select {
		case p, ok := <-cd.C:
			if !ok {
				return
			}
			log.Printf("receive packet is empty?: %v\n", p.IsEmpty())
			log.Printf("upload transmission timestamp count is %d", len(p.TransmissionTimestamps))
			log.Printf("and timestamp is %s\n", time.Unix(p.TransmissionTimestamps[0].Timestamp, 0))
			by, _ := codec.Encode(&constant.Packet{
				Header: constant.Header{
					SerialId:  p.Header.SerialId,
					Version:   p.Header.Version,
					Timestamp: p.Header.Timestamp,
					Source:    p.Header.Target,
					Target:    p.Header.Source,
				},
				Action: constant.AckAction,
			})
			conn.Write(by)
		}
	}
}
