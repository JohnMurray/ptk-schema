NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

default: all

all: clean build test

build: 
	@echo "$(OK_COLOR)==> Building $(NO_COLOR)"

	echo "package main" > version.go
	echo "" >> version.go
	echo "var Version string = \"`cat .version`\"" >> version.go

	mkdir -p ./bin
	go build -o ./bin/schema *.go

clean:
	@echo "$(OK_COLOR)==> Cleaning$(NO_COLOR)"
	rm -rf ./bin/
	rm -rf version.go

test:
	@echo "$(OK_COLOR)==> Testing Schema...$(NO_COLOR)"
	go test ./...

todo:
	@echo "TODOs: "
	@grep -Rni 'todo: ' .
