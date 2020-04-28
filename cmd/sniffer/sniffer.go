package main

import (
	"fmt"
	"io"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func tlsDump(src gopacket.PacketDataSource, dst io.Writer) error {
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
			if layerType == layers.LayerTypeTLS && len(tls.Handshake) > 0 {
				_, err := fmt.Fprintf(dst, "%s,%s,%s,%s,%d\n",
					ip4.SrcIP, tcp.SrcPort, ip4.DstIP, tcp.DstPort, len(tcp.Options))
				if err != nil {
					return fmt.Errorf("error write to destination: %w", err)
				}
			}
		}
	}

	return nil
}
