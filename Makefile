all: count numberize randomize

count: .bin/count

.bin/count: $(shell find ./cmd/count ./internal/app/count ./internal/pkg -type f -name '*.go')
	go build -o .bin/count ./cmd/count

numberize: .bin/numberize

.bin/numberize: $(shell find ./cmd/numberize ./internal/app/numberize ./internal/pkg -type f -name '*.go')
	go build -o .bin/numberize ./cmd/numberize

randomize: .bin/randomize

.bin/randomize: $(shell find ./cmd/randomize ./internal/app/randomize ./internal/pkg -type f -name '*.go')
	go build -o .bin/randomize ./cmd/randomize

.PHONY: deps
deps:
	dep ensure

.PHONY: test
test:
	go test -v ./...

.PHONY: install
install: 
	go install ./...

.PHONY: uninstall
uninstall:
	go clean -i ./cmd/count
	go clean -i ./cmd/numberize
	go clean -i ./cmd/randomize

.PHONY: clean
clean:
	rm -rf .bin/

# vim: noexpandtab
