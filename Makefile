#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
DOCKER := $(shell which docker)

build:
	go mod tidy \
	&& go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o hobbyd cmd/hobbyd/main.go
.PHONY: build
BINS+=hobbyd

install:build
	cp -f hobbyd ${GOPATH}/bin && ln -nsf ${GOPATH}/bin/hobbyd ${GOPATH}/bin/coeusd

init:
	ignite chain init --skip-proto -y

ignite-build:
	ignite chain build -y --debug --clear-cache --check-dependencies -v

# legacy version 0.11.6
protoVer=0.13.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace --user 0 $(protoImageName)

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh
.PHONY: proto-gen

debug:
	hobbyd start --pruning=nothing --evm.tracer=json --log_level trace \
                 --json-rpc.api eth,txpool,personal,net,debug,web3,miner \
                 --api.enable --json-rpc.enable --json-rpc.address 0.0.0.0:8545 \
                 --json-rpc.ws-address 0.0.0.0:8546

start:
	hobbyd start --pruning=nothing --json-rpc.api eth,txpool,personal,net,debug,web3,miner \
                 --api.enable --json-rpc.enable --json-rpc.address 0.0.0.0:8545  --json-rpc.ws-address 0.0.0.0:8546

serve: install start

docker: clean
	docker build --tag coeus-node -f ./Dockerfile .

reset: install init start

docker-test: install
	docker build --tag coeus-node -f ./Dockerfile.test .

clean:
	rm -rf $(BINS) $(CLEAN)

