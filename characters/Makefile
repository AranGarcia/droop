proto:
	protoc \
	--go_out=internal/ \
	--go_opt=paths=import \
	--go_opt=module=github.com/AranGarcia/droop/characters/internal \
	--go-grpc_out=internal/ \
	--go-grpc_opt=paths=import \
	--go-grpc_opt=module=github.com/AranGarcia/droop/characters/internal \
	internal/adapters/primary/grpc/proto/service.proto

clean:
	rm -rf internal/adapters/primary/grpc/genproto