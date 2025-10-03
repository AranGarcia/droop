module github.com/AranGarcia/droop/characters

go 1.24.6

replace (
	github.com/AranGarcia/droop/protoapis => ../../protoapis
	github.com/AranGarcia/droop/shared/common-errors => ../../shared/common-errors
	github.com/AranGarcia/droop/shared/mongotools => ../../shared/mongotools
)

require (
	github.com/AranGarcia/droop/protoapis v0.0.0-00010101000000-000000000000
	github.com/AranGarcia/droop/shared/common-errors v0.0.0-00010101000000-000000000000
	github.com/AranGarcia/droop/shared/mongotools v0.0.0-00010101000000-000000000000
	github.com/go-playground/validator/v10 v10.22.0
	github.com/google/go-cmp v0.7.0
	go.mongodb.org/mongo-driver v1.17.1
	google.golang.org/grpc v1.75.1
)

require (
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
