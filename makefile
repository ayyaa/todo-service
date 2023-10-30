

.PHONY: all init generate generate_mocks

all: build/main

run:
	go run cmd/main.go

build/main: cmd/main.go
	@echo "Building..."
	go build -o $@ $<

init: 
	go mod tidy

test:
	go test -short -coverprofile coverage.out -v ./...

generate: generate_mocks

INTERFACES_GO_FILES := $(shell find repository -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))