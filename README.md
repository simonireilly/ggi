# About

This project is a golang cli for generating gitignore files.

- It utilizes a [fuzzy search](https://github.com/ktr0731/go-fuzzyfinder),
- A full list of gitignore from [GitHubs gitignore files](https://github.com/github/gitignore)
- A github synchronization client [golang GitHub/Google synchronization client](https://github.com/google/go-github)

- [About](#about)
  - [Install](#install)
  - [Specification](#specification)
  - [Design](#design)

## Install



## Specification

1. We have a command go-gitignore-it that starts a session
2. We can search through all the files available at `github/gitignore`
3. We can combine multiple `gitiignore` files
4. We can add our own gitignore files
5. We can install globally ad run anywhere on the file system

## Design

1. A main function for running at the cli:
2. A function for pulling gitignores and storing them:
   1. Includes versioning
   2. Stores locally
   3. Works with latest when offline
   4. cli ships with latest version and cli update command checks and pulls new versions
3. A function for writing gitignore files:
   1. It will mod/alert if existing `.gitignore` is present
   2. It can be run in append only
