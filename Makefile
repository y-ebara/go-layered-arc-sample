DB_URL=postgresql://root:passwd@localhost:5432/test_app2?sslmode=disable
MOCK_FILES := "internal/service/todo.go"

sqlc-generate:
	sqlc generate

migrateup:
	migrate -path db/migrations/ -database "$(DB_URL)"

migratedown:
	migrate -path db/migrations/-database "$(DB_URL)"

test:
	go test -v -cover  ./...

mock-generate:
	@for file in ${MOCK_FILES}; do mockgen -destination "mock/mocks.go" -source "$$file"; done
