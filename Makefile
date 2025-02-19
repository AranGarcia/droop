proto: proto-dir
	buf generate
	cd proto/gen; go mod tidy

proto-dir:
	mkdir -p proto/gen
	cd proto/gen; go mod init github.com/AranGarcia/droop/proto/gen

test:
	go test ./services/characters/...
	go test ./services/dnd/...
	go test ./services/initiatives/...

clean:
	rm -rf proto/gen