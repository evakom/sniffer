package main

import (
	"errors"
	"log"
	"testing"

	"github.com/google/gopacket/pcap"
)

const (
	pcapTestFile = "sniffer_test.pcap"
	testData     = `192.168.137.7,60208,3.123.217.208,443(https),3
3.123.217.208,443(https),192.168.137.7,60208,3
192.168.137.7,60208,3.123.217.208,443(https),3
`
)

type testDst struct {
	data []byte
}

func newTestDst() *testDst {
	return &testDst{data: []byte{}}
}

func (t *testDst) Write(p []byte) (int, error) {
	if len(t.data) >= len(testData) {
		return 0, errors.New("dst write test error")
	}

	t.data = append(t.data, p...)

	return len(p), nil
}

func (t testDst) String() string {
	return string(t.data)
}

func TestTlsDump(t *testing.T) {
	tests := []struct {
		name         string
		testDataFile string
		dst          *testDst
		wantDst      string
		wantErr      bool
	}{
		{
			name:         "Simple capture",
			testDataFile: pcapTestFile,
			wantDst:      testData,
		},
		{
			name:         "Error write dst",
			testDataFile: pcapTestFile,
			wantDst:      testData,
			wantErr:      true,
		},
	}

	dst := newTestDst()

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handle, err := pcap.OpenOffline(tt.testDataFile)
			if err != nil {
				log.Fatal("could not open test file stream:", err)
			}

			err = tlsDump(handle, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("tlsDump() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if gotDst := dst.String(); gotDst != tt.wantDst {
				t.Errorf("tlsDump() gotDst = %v, want = %v", gotDst, tt.wantDst)
			}
		})
	}
}
