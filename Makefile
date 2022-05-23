.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...
