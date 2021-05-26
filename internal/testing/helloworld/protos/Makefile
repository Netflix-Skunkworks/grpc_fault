all: protogen

GOBIN_TOOL	= $(shell which gobin || echo $(GOBIN)/gobin)

.PHONY: protogen
protogen: helloworld.pb.go helloworld_grpc.pb.go

helloworld.pb.go helloworld_grpc.pb.go: helloworld.proto Makefile
	protoc --plugin=protoc-gen-titusgo=$(shell $(GOBIN_TOOL) -p github.com/golang/protobuf/protoc-gen-go@v1.4.3) \
			--titusgo_out=plugins=grpc:. \
			helloworld.proto
