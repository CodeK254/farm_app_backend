build:
	@go build -o bin/farm_app_backend main.go

test:
	@go test -v ./...

run: build
	@./bin/farm_app_backend
