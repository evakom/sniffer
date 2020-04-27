package main

import (
	"log"
	"time"

	"github.com/google/gopacket/dumpcommand"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/pcap"
)

func main() {
	defer util.Run()()

	handle, err := pcap.OpenLive("enp5s0", 1024, false, 30*time.Second)
	if err != nil {
		log.Fatalf("could not open live stream: %v", err)
	}

	dumpcommand.Run(handle)
}
