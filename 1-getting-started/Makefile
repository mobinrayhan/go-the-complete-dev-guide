N := 1
RN := 1

# ALL SERVER RELATED COMMANDS
go-build:
	go build -o build/main ./cmd/server/main.go

# ALL MIGRATIONS COMMANDS
migrate-new:
	 migrate create -ext sql -dir ./internals/db/migrations -seq $(name)

migrate-up:
	go run ./cmd/migrate/main.go -up

migrate-down:
	go run ./cmd/migrate/main.go -down $(N)

migrate-reset:
	go run ./cmd/migrate/main.go -reset $(RN)

migrate-version:
	go run ./cmd/migrate/main.go -version

# ALL SEED COMMANDS
seed-dev:
	APP_ENV=development go run ./cmd/seed/main.go

seed-prod:
	APP_ENV=production  go run ./cmd/seed/main.go