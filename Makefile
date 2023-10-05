test:
	go test ./...

golint:
	staticcheck ./...

migrateupall:
	migrate -path ../db/migration -database "postgresql://root:secret@localhost:5432/test_db?sslmode=disable" -verbose up

.PHONY: test, golint, migrateupall