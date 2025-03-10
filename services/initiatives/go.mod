module github.com/AranGarcia/droop/initiatives

go 1.24.0

replace (
	github.com/AranGarcia/droop/proto/gen => ../../proto/gen
	github.com/AranGarcia/droop/shared => ../../shared
)

require (
	github.com/AranGarcia/droop/proto/gen v0.0.0-00010101000000-000000000000
	github.com/AranGarcia/droop/shared/redistools v0.0.0-20241105061033-22dbd0ad8d15
	github.com/redis/go-redis/v9 v9.7.0
	google.golang.org/grpc v1.71.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
