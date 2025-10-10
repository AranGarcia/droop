test:
	go test ./services/characters/...
	go test ./services/dnd/...
	go test ./services/initiatives/...

run-dbs:
	docker compose up mongodb redis kafka

run-services:
	docker compose up characters dnd