default: all


all: build


build:
	go build -o bin/schema main.go config.go
