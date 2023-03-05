.PHONY: deps
deps:
	go install github.com/golang/mock/mockgen@v1.6.0
	git submodule update --remote

.PHONY: generate
generate:
	go generate ./internal/adapters/core/files/files.go

.PHONY: mocks
mocks:
	mockgen -source=internal/ports/core.go -destination=mocks/ports/core.go

.PHONY: build
build: generate
	go build -o ggi cmd/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: local-install
local-install: build
	sudo mv ggi /usr/local/bin/ggi

.PHONY: ci
ci: deps generate mocks test build
