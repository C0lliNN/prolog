# Prolog
A Logging service created while reading the book 'Distributed Services with Go' from Travis Jeffery.

## Installation
* Go 1.17+ necessary
* [Install Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)
* [Install Kubernetes CLI](https://kubernetes.io/docs/tasks/tools/)
* [Install Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
* Run `go get ./...`
* Run `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
* Run `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
* Run `go get github.com/cloudflare/cfssl/cmd/cfssl`
* Run `go get github.com/cloudflare/cfssl/cmd/cfssljson`
* Run `make init`
* Run `make gencert`

## How to run tests
```bash
make test
```

## How to run it locally
TODO
