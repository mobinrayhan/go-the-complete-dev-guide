migrate-up:
	go run cmd/migrate/main.go -up
	
migrate-down:
	go run cmd/migrate/main.go -down 1

migrate-force:
	MIGRATE_VERSION ?= -1;
	go run cmd/migrate/main.go -force $(MIGRATE_VERSION)

migrate-version:
	go run cmd/migrate/main.go -version

migrate-new:
	migrate create -ext sql -dir ./internals/db/migrations -seq $(name)
