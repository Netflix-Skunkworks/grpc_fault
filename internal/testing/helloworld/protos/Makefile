all: protogen

GOBIN_TOOL	= $(shell which gobin || echo $(GOBIN)/gobin)

.PHONY: protogen
protogen: helloworld.pb.go helloworld_grpc.pb.go

helloworld.pb.go helloworld_grpc.pb.go: helloworld.proto
	protoc \
		--plugin=protoc-gen-go=$(shell $(GOBIN_TOOL) -p google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0) \
		--plugin=protoc-gen-go-grpc=$(shell $(GOBIN_TOOL) -p google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1) \
		--go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        helloworld.proto