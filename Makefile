.PHONY:frontend
GO           ?= go
GOFMT        ?= $(GO)fmt
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))

pkgs          = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
DIRNAME                 ?= $(shell dirname $(shell pwd))

#TAG                     ?= $(shell date +%s)
TAG                     ?= $(shell git rev-parse --short HEAD)


RUN_ENV                 ?= test

style:
	@echo ">> checking code style"
	@! $(GOFMT) -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build:
	@echo ">> go build ..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags -w -o xdhuxc-message main.go


clean:
	@echo ">> remove xdhuxc-message"
	@rm xdhuxc-message

frontend:
	@cd frontend && npm install --registry=https://registry.npm.taobao.org
	@cd frontend && npm run build
