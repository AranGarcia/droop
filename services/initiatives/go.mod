module github.com/AranGarcia/droop/initiatives

go 1.23.2

replace github.com/AranGarcia/droop/proto/gen => ../../proto/gen

require (
	github.com/AranGarcia/droop/proto/gen v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.7.0
	google.golang.org/grpc v1.67.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)