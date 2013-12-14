default: all

all: build test

build:
	mkdir ./bin
	go build -o ./bin/schema *.go

test:
	go test
