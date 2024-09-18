postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres

postgres:
	docker exec -it postgres15 psql --username=postgres

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5433/go-chat?sslmode=disable" -verbose down

serve:
	go run cmd/main.go

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown serve