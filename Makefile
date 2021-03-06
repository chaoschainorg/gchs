# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gchs android ios gchs-cross evm all clean
.PHONY: gchs-linux gchs-linux-386 gchs-linux-amd64 gchs-linux-mips64 gchs-linux-mips64le
.PHONY: gchs-linux-arm gchs-linux-arm-5 gchs-linux-arm-6 gchs-linux-arm-7 gchs-linux-arm64
.PHONY: gchs-darwin gchs-darwin-386 gchs-darwin-amd64
.PHONY: gchs-windows gchs-windows-386 gchs-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

gchs:
	$(GORUN) build/ci.go install ./cmd/gchs
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gchs\" to launch gchs."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gchs.aar\" to use the library."

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gchs.framework\" to use the library."

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

gchs-cross: geth-linux geth-darwin geth-windows geth-android geth-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gchs-*

gchs-linux: geth-linux-386 geth-linux-amd64 geth-linux-arm geth-linux-mips64 geth-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-*

gchs-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gchs
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep 386

gchs-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gchs
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep amd64

gchs-linux-arm: geth-linux-arm-5 geth-linux-arm-6 geth-linux-arm-7 geth-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep arm

gchs-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gchs
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep arm-5

gchs-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gchs
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep arm-6

gchs-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gchs
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep arm-7

gchs-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gchs
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep arm64

gchs-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gchs
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep mips

gchs-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gchs
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep mipsle

gchs-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gchs
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep mips64

gchs-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gchs
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gchs-linux-* | grep mips64le

gchs-darwin: geth-darwin-386 geth-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gchs-darwin-*

gchs-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gchs
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-darwin-* | grep 386

gchs-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gchs
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-darwin-* | grep amd64

gchs-windows: geth-windows-386 geth-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gchs-windows-*

gchs-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gchs
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-windows-* | grep 386

gchs-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gchs
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gchs-windows-* | grep amd64
