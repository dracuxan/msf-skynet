.PHONY: all run build clean

deps:
	@go mod tidy

run: build
	@./bin/msf-skynet

build:
	@go build -o bin/msf-skynet

clean:
	@rm -rf bin
