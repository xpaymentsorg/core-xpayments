.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/xpaymentsorg/go-xpayments/cmd/swarm
	go build -o bin/XPS     github.com/xpaymentsorg/go-xpayments/cmd/XPS
	go build -o bin/bootnode github.com/xpaymentsorg/go-xpayments/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/xpaymentsorg/go-xpayments/swarm/...
