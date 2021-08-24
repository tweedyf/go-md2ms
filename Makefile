GO111MODULE ?= on

export GO111MODULE

.PHONY:
build: bin/go-md2man

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: test
test:
	@go test $(TEST_FLAGS) ./...

bin/go-md2man: go.mod go.sum md2man/* *.go
	@mkdir -p bin
	CGO_ENABLED=0 go build $(BUILD_FLAGS) -o $@

.PHONY: mod
mod:
	@go mod tidy

.PHONY: check-mod
check-mod: # verifies that module changes for go.mod and go.sum are checked in
	@hack/ci/check_mods.sh

.PHONY: vendor
vendor: mod
	@go mod vendor -v

