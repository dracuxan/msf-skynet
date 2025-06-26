.PHONY: all run build clean

deps:
	@go mod tidy

run: build
	@./bin/cli

build:
	@go build -o bin/cli

clean:
	@rm -rf bin
