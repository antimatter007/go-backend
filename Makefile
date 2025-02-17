# Makefile

# Database configuration using environment variables with default values
DB_USER ?= postgres
DB_PASSWORD ?= admin
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_NAME ?= go_backend
DB_SSLMODE ?= disable
DB_URL = postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Docker network
network:
	docker network create bank-network

# PostgreSQL container
postgres:
	docker run --name postgres --network bank-network -p 5432:5432 \
	-e POSTGRES_USER=$(DB_USER) \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-d postgres:14-alpine

# MySQL container (if needed)
mysql:
	docker run --name mysql8 -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=admin \
	-d mysql:8

# Create database
createdb:
	docker exec -it postgres createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

# Drop database
dropdb:
	docker exec -it postgres dropdb $(DB_NAME)

# Migrations
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

# Create a new migration
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

# Generate database documentation
db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

# Generate SQLC code
sqlc:
	sqlc generate

# Run tests
test:
	go test -v -cover -short ./...

# Run the server
server:
	go run main.go

# Generate mocks
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/antimatter007/go-backend/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/antimatter007/go-backend/worker TaskDistributor

# Generate protobuf files
proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

# Evans CLI for gRPC testing
evans:
	evans --host localhost --port 9090 -r repl

# Redis container
redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

# Phony targets
.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis
