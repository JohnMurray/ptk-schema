default: all

all: build test

create-version-file:
	echo "package main" > version.go
	echo "" >> version.go
	echo "var Version string = \"`cat .version`\"" >> version.go

build: create-version-file
	mkdir -p ./bin
	go build -o ./bin/schema *.go

test:
	go test

todo:
	@echo "TODOs: "
	@grep -Rni 'todo: ' .
