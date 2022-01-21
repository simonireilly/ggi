ignores:
	rm -rf ./gitignore
	git clone git@github.com:github/gitignore.git --depth 1
	rm -rf ./gitignore/.git
	rm -rf ./gitignore/.github

build:
	CGO_ENABLED=0 go build -trimpath $(LDFLAGS) -o $@
