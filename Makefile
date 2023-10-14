DB_URL=postgresql://root:secret@localhost:5432/kanban?sslmode=disable
CONTAINER_TOOL=podman
postgres:
	"$(CONTAINER_TOOL)" run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb:
	"$(CONTAINER_TOOL)" exec -it postgres createdb --username=root --owner=root kanban
dropdb:
	"$(CONTAINER_TOOL)" exec -it postgres dropdb --username=root kanban
migrateup:
	script/migrate -path sql/schema -database "$(DB_URL)" -verbose up
migratedown:
	script/migrate -path sql/schema -database "$(DB_URL)" -verbose down
sqlc:
	script/sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
