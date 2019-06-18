# About

This project is a golang cli for generating gitignore files.

It utilizes a [fuzzy search](https://github.com/ktr0731/go-fuzzyfinder), [GitHubs gitignore files](https://github.com/github/gitignore) and [golang GitHub/Google synchronization client](https://github.com/google/go-github)

- [About](#About)
  - [Specification](#Specification)
  - [Design](#Design)

## Specification

1. We have a command go-gitignore-it that starts a session
2. We can search through all the files available at `github/gitignore`
3. We can combine multiple `gitiignore` files
4. We can add our own gitignore files
5. We can install globally ad run anywhere on the file system

## Design

1. A main function for running at the cli
2. A function for pulling gitignores and storing them
   1. Includes versioning
   2. Stores locally
   3. Works with latest when offline
   4. cli ships with latest version and cli update command checks and pulls new versions
