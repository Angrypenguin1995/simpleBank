createpostgres_o:
	docker run --name postgres12 -p 54000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createpostgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

start_postgres12:
	docker start postgres12

start_postgres12_and_connect:
	docker start postgres12 /bin/sh

stop_postgres12:
	docker stop postgres12

simple_bank_migrate_up:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up

simple_bank_migrate_up_one_migration:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

simple_bank_migrate_down:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down

simple_bank_migrate_down_one_migration:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

simple_bank_migrate_up_o:
	migrate -path db/migration -database "postgresql://root:password@localhost:54000/simple_bank?sslmode=disable" --verbose up

simple_bank_migrate_down_o:
	migrate -path db/migration -database "postgresql://root:password@localhost:54000/simple_bank?sslmode=disable" --verbose down

sqlc_generate:
	sqlc generate

run_test:
	go test -v -cover ./...

start_server:
	go run main.go

mockdb_storego:
	mockgen -package mockdb -destination db/mock/store.go github.com/angrypenguin1995/simple__bank/db/sqlc Store

.PHONY: createpostgres createdb dropdb