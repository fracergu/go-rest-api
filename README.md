# go-rest-api

## Description

This is a simple REST API written in Go, just for practice.

## Usage

Clone the repository and run the following command:

```bash
go install
```

Then, run the following command to start the server:

```bash
go run main.go
```

## Swaggo

To generate the Swagger documentation, install swaggo on `$GOPATH` with the following command:

```bash
go get github.com/swaggo/swag/cmd/swag
```

Then, run the following command to generate the documentation:

```bash
swag init
```

If `$GOPATH/bin` is not in your `$PATH`, remeber to add it yo your `$PATH`:

```bash
# Add the following line to your .bashrc or .zshrc
export PATH=$PATH:$(go env GOPATH)/bin
```
