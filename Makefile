start:
	migrate create -ext sql -dir db/migration -seq init_schema
postgres:
	docker run --name traintickets -p 5434:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it traintickets createdb --username=root --owner=root traintickets

dropdb: 
	docker exec -it traintickets dropdb traintickets

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/traintickets?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/traintickets?sslmode=disable" -verbose down 

sqlc:
	docker run --rm -v /mnt/c/Users/denga/OneDrive/Desktop/readygo/traintickets:/src -w /src kjconroy/sqlc generate

test: 
	go test -v -cover ./...

.PHONY:
	start postgres createdb dropdb migrateup migratedown sqlc test
	