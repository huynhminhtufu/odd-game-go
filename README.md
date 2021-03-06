# OddGame.io Game Server

## Prerequisites

Go: https://golang.org/doc/install
Protoc: https://github.com/golang/protobuf

## Development

- Install [Go](https://golang.org/) and [Docker](https://docs.docker.com/install/).
- Install [gRPC](https://grpc.io/docs/quickstart/go/) `go get google.golang.org/grpc@v1.28.1`
- Install protobuf for MacOS `brew install protobuf`
- Run `docker-compose up mongo -d` to start MongoDB in background (or you can start with local MongoDB).
- Run `make depend` to install dependencies
- Run `go run ./cmd/server/main.go` to start development.

## Production

- Install [Docker](https://docs.docker.com/install/).
- Run `make build` to build.
- Run `make start` to start.
- Run `make stop` to stop.

## Workflow

### Branch naming:

- Feature: feature/add-new-button
- Hotfix: hotfix/fix-bug-abc
- Improvement: improve/improve-button-abc

### Commits:

- Commit should not capitalize first character, example: migrate to hooks

## Stack:

- Go
- MongoDB
- Docker
