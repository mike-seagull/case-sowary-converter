BINARY_NAME=csc

build:
	go build -o ./bin/${BINARY_NAME} ./cmd/${BINARY_NAME}

test:
	go test -v pkg/case_converter

clean:
	go clean
	rm -f ./bin/${BINARY_NAME}

all: clean test build