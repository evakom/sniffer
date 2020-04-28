package main

import (
	"log"
	"strings"
	"testing"

	"github.com/google/gopacket/pcap"
)

const pcapTestFile = "test.pcap"

func TestTlsDump(t *testing.T) {
	handle, err := pcap.OpenOffline(pcapTestFile)
	if err != nil {
		log.Fatal("could not open test file stream:", err)
	}

	tests := []struct {
		name    string
		wantDst string
		wantErr bool
	}{
		{
			name: "Simple capture",
			wantDst: `
192.168.137.7,60208,3.123.217.208,443(https),3
3.123.217.208,443(https),192.168.137.7,60208,3
192.168.137.7,60208,3.123.217.208,443(https),3
`,
		},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := new(strings.Builder)

			err := tlsDump(handle, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("tlsDump() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if gotDst := "\n" + dst.String(); gotDst != tt.wantDst {
				t.Errorf("tlsDump() gotDst = %v, want = %v", gotDst, tt.wantDst)
			}
		})
	}
}
