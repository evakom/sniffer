package main

import (
	"fmt"
	"log"
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

	fmt.Println("Starting to read packets...")

	tlsDump(handle)
}

func tlsDump(src gopacket.PacketDataSource) {
	var (
		eth     layers.Ethernet
		ip4     layers.IPv4
		tcp     layers.TCP
		tls     layers.TLS
		decoded []gopacket.LayerType
	)

	source := gopacket.NewPacketSource(src, layers.LayerTypeEthernet)
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &tcp, &tls)

	for packet := range source.Packets() {
		if err := parser.DecodeLayers(packet.Data(), &decoded); err != nil {
			// err contains error parse layer
			continue
		}

		for _, layerType := range decoded {
			if layerType == layers.LayerTypeTLS {
				fmt.Println("    TLS: ", ip4.SrcIP, "->", ip4.DstIP, tcp.SrcPort, "->", tcp.DstPort, len(tcp.Options))
				//fmt.Println(tls.Handshake[0].Version)
			}
		}
	}
}
