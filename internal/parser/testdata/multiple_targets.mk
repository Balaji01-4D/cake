.PHONY: all build test clean

all: build test

## Compile the binary for production
build: main.go utils.go
	go build -o bin/app .

test: main.go
	go test -v ./...

## Clean the build directory
clean:
	rm -rf bin/
	rm -f *.o