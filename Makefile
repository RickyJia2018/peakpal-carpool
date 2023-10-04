DB_URL=postgresql://root:secret@localhost:5432/peakpal_carpool_db?sslmode=disable

createdb:
	docker exec -it peakpal_db_container createdb --username=root --owner=root peakpal_carpool_db

createTestdb:
	docker exec -it peakpal_db_container createdb --username=root --owner=root peakpal_carpool_test

dropdb:
	docker exec -it peakpal_db_container dropdb peakpal_carpool_db

dropTestdb:
	docker exec -it peakpal_db_container dropdb peakpal_carpool_test

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/rickyjia2018/peak-pal-server/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/rickyjia2018/peak-pal-server/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f docs/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=peakpal_carpool \
	proto/*.proto
	statik -src=./docs/swagger -dest=./docs

evans:
	evans --host localhost --port 9091 -r repl
server:
	go run main.go

.PHONY: network postgres createdb createTestdb dropdb dropTestdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans 
