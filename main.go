package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	// nolint:gomnd
	handle, err := pcap.OpenLive("enp5s0", 2048, false, 30*time.Second)
	if err != nil {
		log.Fatalf("could not open live stream: %v", err)
	}

	tcpDump(handle)
}

func tcpDump(src gopacket.PacketDataSource) {
	source := gopacket.NewPacketSource(src, layers.LayerTypeEthernet)
	//source.Lazy = true
	source.NoCopy = true
	//source.DecodeStreamsAsDatagrams = true
	_, _ = fmt.Fprintln(os.Stderr, "Starting to read packets...")

	var count int

	var bytes int64

	for packet := range source.Packets() {
		count++

		bytes += int64(len(packet.Data()))
		fmt.Println(packet)
	}
}
