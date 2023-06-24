# Go gRPC Example

## Description

This is a simple example of a gRPC server and client written in Go.

## Prerequisites
[Protocol Buffers](https://protobuf.dev/)
[Protocol Buffers Plugin for Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
[Golang](https://golang.org/)

### Arch Linux
```
sudo pacman -S protobuf

sudo pacman -S go
```

## Check protobuf version
```
protoc --version
```
Must be 3.0.0 or higher

## Protocol Buffers Plugin for Go
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Export PATH
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Generate gRPC code
```
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

## SQL Script for SQLite3
```
create table category (
    id string,
    name string,
    description string,
);
```