.PHONY: deps
deps:
	go install github.com/gobuffalo/packr/v2/packr2@v2.8.3

.PHONY: build
build:
	packr2 build

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	packr2 clean
