.PHONY: deps
deps:
	go install github.com/gobuffalo/packr/v2/packr2@v2.8.3
	git submodule update --remote

.PHONY: generate
generate:
	go generate ./internal/adapters/core/files/files.go


.PHONY: build
build: generate
	packr2 build -o ggi/ ./cmd

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	packr2 clean

.PHONY: local-install
local-install: build
	sudo mv ggi /usr/local/bin/ggi
