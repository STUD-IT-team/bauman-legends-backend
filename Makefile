
migration-create-sql:
	goose -dir=./migrations create $(NAME) sql

swag-generate-app:
	swag init --parseDependency --parseInternal  -g internal/ports/handlers/feed.go -o cmd/api/docs/ -g cmd/api/main.go