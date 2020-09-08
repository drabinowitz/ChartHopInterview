.PHONY: all build start

all: install build start
install:
	go mod tidy
start:
	docker start local-dynamo-test
	./build/server
build:
	go build -o build/server
