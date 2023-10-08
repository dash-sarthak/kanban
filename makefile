migrateup:
    migrate -path sql/schema -database "postgresql://user:pass@localhost:5432/kanban?sslmode=disable" -verbose up

migratedown:
    migrate -path sql/schema -database "postgresql://user:pass@localhost:5432/kanban?sslmode=disable" -verbose down

.PHONY: migrateup migratedown

