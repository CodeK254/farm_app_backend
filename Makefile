build:
	@go build -o bin/farm_app_backend main.go

test:
	@go test -v ./...

run: build
	@./bin/farm_app_backend

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

migrate-force:
	@go run cmd/migrate/main.go $(filter-out $@,$(MAKECMDGOALS)) force