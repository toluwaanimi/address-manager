
PROTO_SRC_DIR := ${PWD}/proto
PROTO_DST_DIR := ${PWD}/pb

.PHONY: proto
proto:
	go install github.com/golang/protobuf/protoc-gen-go
	mkdir -p ${PROTO_DST_DIR} && protoc -I=${PROTO_SRC_DIR} --go_out=plugins=grpc:${PROTO_DST_DIR} ${PROTO_SRC_DIR}/*.proto


clean:
	rm pb/*

run:
	go run main.go


gen-mocks:
	go generate ./...

local: proto mocks
	go run main.go