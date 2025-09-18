MIGRATE_VERSION ?= -1;
ENV ?= development;

migrate-up:
	go run cmd/migrate/main.go -up
	
migrate-down:
	go run cmd/migrate/main.go -down 1

migrate-force:
	go run cmd/migrate/main.go -force $(MIGRATE_VERSION)

migrate-version:
	go run cmd/migrate/main.go -version

migrate-new:
	migrate create -ext sql -dir ./internals/db/migrations -seq $(name)
	
seed:
	APP_ENV=$(ENV) go run cmd/migrate/main.go -seed

build:
	go build -o ./build/main ./cmd/server

run-prod:
	./build/main -mode=production
	
run-dev:
	air -c .air.toml