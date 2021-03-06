.PHONY: build run test docker-build docker-run

IFACE ?= enp5s0
PROM ?= false
SIZE ?= 2048
HTTP ?= no

build:
	go build -o app/sniffer cmd/sniffer/*.go

run:
	./app/sniffer -i $(IFACE) -p=$(PROM) -s $(SIZE) -h $(HTTP)

test:
	go test ./... -v -cover

docker-build:
	docker build -t sniffer -f ./build/package/sniffer/Dockerfile .

docker-run:
	docker run --network host sniffer /root/app/sniffer -i $(IFACE) -p=$(PROM) -s $(SIZE) -h $(HTTP)
