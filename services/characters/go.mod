module github.com/AranGarcia/droop/characters

go 1.23.2

replace (
	github.com/AranGarcia/droop/proto/gen => ../../proto/gen
	github.com/AranGarcia/droop/shared/mongotools => ../../shared/mongotools
)

require (
	github.com/AranGarcia/droop/proto/gen v0.0.0-00010101000000-000000000000
	github.com/AranGarcia/droop/shared/mongotools v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/go-cmp v0.6.0
	go.mongodb.org/mongo-driver v1.17.1
	google.golang.org/grpc v1.67.1
)

require (
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)
