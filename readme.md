# Template for contract binding service using Go

## TOC

- OpenAPI for service [api/contractbind/v1/api.swagger.json](./api/contractbind/v1/api.swagger.json)
- entrypoint for go server [cmd/server/main.go](./cmd/server/main.go)
- docker image [deploy/contractbind.dockerfile](./deploy/contractbind.dockerfile)
- k8s (helm) configuration [deploy](./deploy),[deploy/templates/deployment.yaml](./deploy/templates/deployment.yaml)
- GRPC API implementation [pkg/api/api.go](./pkg/api/api.go)
- Contract bindings for Go [pkg/contracts/commit_reveal.go](./pkg/contracts/commit_reveal.go)
- Contract test using popular Go BDD framework [pkg/contracts/crowdfund_test.go](./pkg/contracts/crowdfund_test.go)
- GRPC/GRPC-HTTP-Gateway protobuf definitions [proto/contractbind/v1/api.proto](./proto/contractbind/v1/api.proto)
- GRPC configurations [buf.gen.yaml](./buf.gen.yaml)
- Taskfile (like makefile) [Taskfile.yaml](./Taskfile.yaml)

### Commands

Before you begin, make sure you have the following prerequisites:

- Task runner installed (https://taskfile.dev/installation/)
- Docker installed to build and package the service.
- Helm (https://helm.sh/) installed for Kubernetes deployment and management.
- kubectl (Kubernetes command-line tool) configured to interact with your Kubernetes cluster.
- Go for development installed (https://go.dev/doc/install)
- Buf for managing protobuf files installed (https://buf.build/docs/installation)

+ `task build` will build application image and chart and store it at ./build
+ `task deploy` will install or upgrade helm deployment, requires to run `task build` first (be sure that values [deploy/values/default.yaml](./deploy/values/default.yaml) are correctly populated, and extended for production enviroment)
+ `task logs` will show logs of running helm deployment, requires to run `task deploy` first
+ `task port-forward` will port forward to service of running helm deployment, reachable at http://127.0.0.1:8080, requires to run `task deploy` first
+ `task clean` will perform clean up, remove helm deployment and build artefacts
+ `task gen` will generate abi and bin for solidity contract, generate Go contract bindings, generate protobuf bindings for go, generate open API files
+ `task dev-mac` will install all dev requirements (macOS/linux-homebrew only supported)

### Testing

Run `go test -timeout 30s -v ./...` for tests on emulated eth backend

To perform integrational test you need to have blockchain running, server running, and then test it via HTTP calls

1. Open three terminal windows in root repository folder
1. Run `task testnet` in first, to run local etherium node API
1. Run `task local-run` in second, to run server at localhost:80
1. Run `go test -timeout 30s -count 1 -run ^TestAPI$ ./testing -v` in third, to run [test](./testing/local_test.go)

## TODO

1. Put your contract into ./contracts

2. https://github.com/ic-n/ethcontractbind/blob/main/proto/contractbind/v1/api.proto#L19
```go
// TODO: define your custom api
```

3. https://github.com/ic-n/ethcontractbind/blob/main/pkg/api/api.go#L54
```go
// TODO: implement API
```

4. https://github.com/ic-n/ethcontractbind/blob/main/pkg/contracts/crowdfund_test.go#L35
```go
// TODO: implement contract test
```

5. https://github.com/ic-n/ethcontractbind/blob/main/testing/local_test.go#L25
```go
// TODO: implement testing make api calls
```

6. Remove TODO section from this file :)
