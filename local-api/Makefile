#!make

SHELL	= bash

ifndef BINARY_NAME
	BINARY_NAME 	:= $(shell basename $(CURDIR))
endif

ifndef NETWORK_NAME
	NETWORK_NAME 	:= $(shell basename $(CURDIR))
endif

ifndef PORT
	PORT 	:= 8080
endif

docker-build:
	docker build -t ${BINARY_NAME} . 

docker-run:
	docker run --rm ${BINARY_NAME}

docker-run-network: docker-build
	docker run \
		--name ${BINARY_NAME} \
		--network=${NETWORK_NAME} \
		-e "PORT=${PORT}" \
		-p ${PORT}:${PORT} \
		--rm ${BINARY_NAME}

test:
	go test -v ./...

# build: test
build:
	CGO_ENABLED=0 GOOS=linux go build -a -o ${BINARY_NAME} .

run: build
	./${BINARY_NAME}