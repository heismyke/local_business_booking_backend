postgres:
		docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine
createdb:
		docker exec -it postgres17 createdb --username=root --owner=root street_business
dropdb:
		docker exec -it postgres17 dropdb street_business
migrateup:
		migrate -path ./internal/db/migration -database "postgresql://root:secret@localhost:5432/street_business?sslmode=disable" --verbose up
migratedown:
	migrate -path ./internal/db/migration -database "postgresql://root:secret@localhost:5432/street_business?sslmode=disable" force 1
	migrate -path ./internal/db/migration -database "postgresql://root:secret@localhost:5432/street_business?sslmode=disable" --verbose down
sqlc:
		sqlc generate
test:
		go test -v -cover ./...
.PHONY: createdb dropdb migrateup migratedown sqlc
