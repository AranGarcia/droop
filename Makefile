proto: proto-dir
	buf generate
	cd proto/gen; go mod tidy

proto-dir:
	mkdir -p proto/gen
	cd proto/gen; go mod init github.com/AranGarcia/droop/proto/gen

test:
	go test ./characters/...
	go test ./dnd/...

clean:
	rm -rf proto/gen