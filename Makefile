run:	
	go run cmd/main.go

tidy:
	go mod tidy

migrate_create:
	migrate create -ext sql -dir internal/migrations/ -seq tables	

migrate_up:
	migrate	-path internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/tsu_toleg?sslmode=disable"	-verbose up

migrate_force:
	migrate -path internal/migrations/ -database "postgresql://postgres:12345@localhost:5432/tsu_toleg?sslmode=disable" force 1

.PHONY: run tidy migrate_create migrate_run migrate_force	