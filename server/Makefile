.PHONY: all build start

all: install build start
install:
	go mod tidy
start:
	./build/server
build:
	go build -o build/server
