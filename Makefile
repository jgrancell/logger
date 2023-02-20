## Copyright 2020-2022 Josh Grancell. All rights reserved.
## Use of this source code is governed by an MIT License
## that can be found in the LICENSE file.
TEST?=$$(go list ./... | grep -v vendor)
WORKDIR=$$(pwd)
BINARY=$$(pwd | xargs basename)
VERSION=$$(grep version main.go | head -n1 | cut -d\" -f2)
GOBIN=${GOPATH}/bin

.PHONY: test

default: build

build:
	go build -o ${BINARY}
	chmod +x ${BINARY}

install: build
	mkdir -p ${GOBIN}
	mv ${BINARY} ${GOPATH}/bin/${BINARY}

## Testing tasks
test:
	rm -f coverage.txt profile.out
	rm -f gosec-report.json
	rm -rf test/test*
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

test-view: test
	go tool cover -html=coverage.txt