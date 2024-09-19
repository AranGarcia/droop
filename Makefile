character-proto:
	protoc \
	--go_out=characters/internal/ \
	--go_opt=paths=import \
	--go_opt=module=github.com/AranGarcia/droop/characters/internal \
	--go-grpc_out=characters/internal/ \
	--go-grpc_opt=paths=import \
	--go-grpc_opt=module=github.com/AranGarcia/droop/characters/internal \
	proto/characters/service.proto

clean:
	rm -rf characters/internal/adapters/primary/grpc/genproto