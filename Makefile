run:	
	go run cmd/main.go

tidy:
	go mod tidy

migrate_create:
	migrate create -ext sql -dir internal/infrastructure/migrations/ -seq tables	

migrate_up:
	migrate	-path internal/infrastructure/migrations/ -database "postgresql://postgres:12345@localhost:5432/tsu_toleg?sslmode=disable"	-verbose up

migrate_force:
	migrate -path internal/infrastructure/migrations/ -database "postgresql://postgres:12345@localhost:5432/tsu_toleg?sslmode=disable" force 1

swag-init:
	swag init --output ./docs --parseDependency --generalInfo ./cmd/main.go

.PHONY: run tidy migrate_create migrate_run migrate_force swag-init