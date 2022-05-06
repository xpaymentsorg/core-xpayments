.PHONY: XPS XPS-cross evm all test clean
.PHONY: XPS-linux XPS-linux-386 XPS-linux-amd64 XPS-linux-mips64 XPS-linux-mips64le
.PHONY: XPS-darwin XPS-darwin-386 XPS-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= latest
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

XPS:
	build/env.sh go run build/ci.go install ./cmd/XPS
	@echo "Done building."
	@echo "Run \"$(GOBIN)/XPS\" to launch XPS."

gc:
	build/env.sh go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	build/env.sh go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	build/env.sh go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

XPS-cross: XPS-linux XPS-darwin
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/XPS-*

XPS-linux: XPS-linux-386 XPS-linux-amd64 XPS-linux-mips64 XPS-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-*

XPS-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/XPS
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-* | grep 386

XPS-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/XPS
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-* | grep amd64

XPS-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/XPS
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-* | grep mips

XPS-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/XPS
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-* | grep mipsle

XPS-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/XPS
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-* | grep mips64

XPS-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/XPS
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/XPS-linux-* | grep mips64le

XPS-darwin: XPS-darwin-386 XPS-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/XPS-darwin-*

XPS-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/XPS
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/XPS-darwin-* | grep 386

XPS-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/XPS
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/XPS-darwin-* | grep amd64

gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
