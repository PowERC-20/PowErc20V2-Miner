LDFLAGS=-ldflags "-w -s"
export GO111MODULE=on
export CGO_ENABLED=0

.PHONY: all
all: release-dirs clean deps build_linux_amd64 build_windows_amd64 build_darwin_amd64 upx

.PHONY: clean
clean:
	- go clean
	- rm -f bin/*

.PHONY: deps
deps:
	go mod tidy

.PHONY: release-dirs
release-dirs:
	mkdir -p bin/

.PHONY: build_windows_amd64
build_windows_amd64:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/${PROJECT_NAME}_${VERSION}_windows_amd64.exe

.PHONY: build_linux_amd64
build_linux_amd64:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/${PROJECT_NAME}_${VERSION}_linux_amd64

.PHONY: build_darwin_amd64
build_darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/${PROJECT_NAME}_${VERSION}_darwin_amd64

.PHONY: upx
upx:
	upx bin/*
