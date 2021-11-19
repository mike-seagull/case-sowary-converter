BINARY_NAME=csc
MAIN_DIR=cmd

all: clean test build

build:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY_NAME}-mac ./${MAIN_DIR}/${BINARY_NAME}
	env GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY_NAME} ./${MAIN_DIR}/${BINARY_NAME}
	env GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY_NAME}-win.exe ./${MAIN_DIR}/${BINARY_NAME}

test:
	go test -v pkg/case_converter

clean:
	go clean
	rm -f ./bin/${BINARY_NAME}*