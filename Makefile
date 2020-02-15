GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	go run main.go

.PHONY: build-mac
build-mac:
	go build -o monkey *.go

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -o monkey

.PHONY: test
test:
	go test -v -race ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
