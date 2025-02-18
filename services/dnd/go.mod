module github.com/AranGarcia/droop/dnd

go 1.24.0

replace github.com/AranGarcia/droop/proto/gen => ../../proto/gen

require (
	github.com/AranGarcia/droop/proto/gen v0.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.47
	go.mongodb.org/mongo-driver v1.17.1
	google.golang.org/grpc v1.70.0
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
