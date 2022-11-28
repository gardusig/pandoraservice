# Go gRPC Service

## Generate models from protobuf

We have defined the `.proto` at `/protos` containing:
- `rpc` call named by `ExampleRequest`, with its:
  - `ClientRequestModel`
  - `ServerResponseModel`

So, how can we generate equivalent models that Go can understand? Just run this command:

```shell
$ protoc --go_out=./go/model/generated protos/example.proto
```