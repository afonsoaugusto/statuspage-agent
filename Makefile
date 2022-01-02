#!make

SHELL	= bash

ifndef BINARY_NAME
	BINARY_NAME 	:= $(shell basename $(CURDIR))
endif

docker-build:
	docker build -t ${BINARY_NAME} . 

docker-run:
	docker run ${BINARY_NAME}

test:
	go test -v ./...

# build: test
build:
	CGO_ENABLED=0 GOOS=linux go build -a -o ${BINARY_NAME} .

run: build
	./${BINARY_NAME}