GOBIN_TOOL		= $(shell which gobin)
GOIMPORT_TOOL	= $(GOBIN_TOOL) -m -run golang.org/x/tools/cmd/goimports@v0.1.0 -w


.PHONY: all
all: helloworld faultinjectorpb

.PHONY: helloworld
helloworld:
	$(MAKE) -C internal/testing/helloworld

.PHONY: faultinjectorpb
faultinjectorpb:
	$(MAKE) -C faultinjector

.PHONY: fmt
fmt:
	$(GOIMPORT_TOOL) $(shell go list -f '{{.Dir}}' ./...)