postgresinit:
	docker run --name pg_chat_db -p 5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=chatdb -d postgres

postgres:
	docker exec -it pg_chat_db psql --username=postgres

createdb:
	docker exec -it pg_chat_db createdb --username=postgres --owner=postgres chatdb

dropdb:
	docker exec -it pg_chat_db dropdb chatdb --username=postgres

.PHONY: postgresinit postgres createdb dropdb
