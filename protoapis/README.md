# Protobuf API repository

Go module with compiled protocol buffer repository for client/server implementations. Since this is used from the `droop` monorepo, other Go modules that implement this must add this as a dependency:

```go.mod
replace (
    github.com/AranGarcia/droop/protoapis => ../../protoapis
)

require (
    github.com/AranGarcia/droop/protoapis latest
)
```

## How to update

Whenever a change is made to a `.proto` file, the protocol buffer code must be re-compiled:
```
buf generate
```