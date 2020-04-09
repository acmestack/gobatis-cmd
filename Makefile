export GO111MODULE=on

build: gobatis-cmd
.PHONY: build

clean:
.PHONY: clean

test:
.PHONY: test

gobatis-cmd: FORCE
    go build -o $@ ./cmd/gobatis-cmd

vendor: go.mod go.sum
    go mod vendor

install:
    go install -o gobatis-cmd ./cmd/gobatis-cmd
.PHONY: install