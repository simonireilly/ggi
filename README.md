# Go Gitignore It

This project is a golang cli for generating gitignore files.

![zsh Terminal using ggi to generate a gitignore](.readme/ggi-in-action.gif)

- [Go Gitignore It](#go-gitignore-it)
  - [Usage](#usage)
  - [Install](#install)
  - [Development](#development)
    - [Dependencies](#dependencies)

## Usage

```bash
ggi <search>
```

## Install

Download the binary for you platform from releases and move the file to your `/usr/local/bin`


## Development

See [Makefile](./Makefile) for commands for testing and building.

### Dependencies

- It utilizes a [fuzzy search](https://github.com/ktr0731/go-fuzzyfinder),
