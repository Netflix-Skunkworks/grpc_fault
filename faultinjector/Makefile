GOBIN_TOOL	= $(shell which gobin || echo $(GOBIN)/gobin)

.PHONY: all
all: protogen

.PHONY: protogen
protogen: faultinjector.pb.go faultinjector_grpc.pb.go

faultinjector.pb.go faultinjector_grpc.pb.go: faultinjector.proto
	protoc \
		--plugin=protoc-gen-go=$(shell $(GOBIN_TOOL) -p google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0) \
		--plugin=protoc-gen-go-grpc=$(shell $(GOBIN_TOOL) -p google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1) \
		--go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        faultinjector.proto

