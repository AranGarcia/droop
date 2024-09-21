proto: characters-proto dnd-proto

characters-proto: proto-dir
	protoc \
	--proto_path=proto \
	--go_out=./proto/gen \
	--go_opt=paths=source_relative \
	--go_opt=Mcharacters/service.proto=proto/characters \
	--go-grpc_out=./proto/gen \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_opt=Mcharacters/service.proto=proto/characters \
	characters/service.proto

dnd-proto: proto-dir
	protoc \
	--proto_path=proto \
	--go_out=./proto/gen \
	--go_opt=paths=source_relative \
	--go_opt=Mdnd/service.proto=proto/dnd \
	--go-grpc_out=./proto/gen \
	--go-grpc_opt=paths=source_relative \
	--go-grpc_opt=Mdnd/service.proto=proto/dnd \
	dnd/service.proto

proto-dir:
	mkdir -p proto/gen

clean:
	rm -rf proto/gen