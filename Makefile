default: all


all: build test


build:
	go build -o bin/schema *.go

test:
	go test
