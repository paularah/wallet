postgres:
	 docker run --name postgres --rm -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest

kill-db:
	docker kill postgres
	 
sqlc:
	sqlc generate

create-db:
	docker exec -it postgres createdb --username=root --owner=root wallet 

drop-db:
	docker exec -it postgres dropdb wallet 

migrate-up: 
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/wallet?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/wallet?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run cmd/main.go  

mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination pkg/db/mock/store.go  github.com/paularah/wallet/pkg/db/sqlc Store 

.PHONY: createdb dropdb postgres migrateup migratedown killdb test mock


dev/init:
	@echo 'setting up dev enviroment'
	make postgres 
	@echo 'Sleeing for 6 seconds'
	sleep 10
	make create-db
	make migrate-up
	make server


.PHONY: dev/init 


