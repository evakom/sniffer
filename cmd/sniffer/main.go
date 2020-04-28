package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	httpserver "github.com/evakom/sniffer/internal/http_server"
	"github.com/google/gopacket/pcap"
)

// If timeout greater than zero it only set the PCAP handle into non-blocking mode.
// If we want to block forever, use zero and libpcap will handle that.
const timeout = 30 * time.Second

func main() {
	iFace := flag.String("i", "eth0", "Interface to read packets from")
	snapLen := flag.Int("s", 65536, "Snap length (number of bytes max to read per packet")
	promisc := flag.Bool("p", false, "Set promiscuous mode")
	listenHTTP := flag.String("h", "", "Host:port for http server listen")

	flag.Parse()

	if *listenHTTP != "" {
		go func() {
			serv := httpserver.New()

			fmt.Println("Starting http server at host:port -", *listenHTTP)
			if err := serv.Start(); err != nil {
				log.Fatal("http server start error", err)
			}
		}()
	}

	handle, err := pcap.OpenLive(*iFace, int32(*snapLen), *promisc, timeout)

	if err != nil {
		log.Fatal("could not open live stream:", err)
	}

	fmt.Println("Starting to read packets ...")

	if err := tlsDump(handle, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
