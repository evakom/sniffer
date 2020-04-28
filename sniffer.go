package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

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
				fmt.Printf("%s:%s -> %s:%s opts:%d\n",
					ip4.SrcIP, tcp.SrcPort, ip4.DstIP, tcp.DstPort, len(tcp.Options))
			}
		}
	}
}
