module github.com/AranGarcia/droop/initiatives

go 1.24.6

replace (
	github.com/AranGarcia/droop/protoapis => ../../protoapis
	github.com/AranGarcia/droop/shared => ../../shared
)

require (
	github.com/AranGarcia/droop/protoapis v0.0.0-00010101000000-000000000000
	github.com/AranGarcia/droop/shared/redistools v0.0.0-20241105061033-22dbd0ad8d15
	github.com/redis/go-redis/v9 v9.7.0
	google.golang.org/grpc v1.75.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
