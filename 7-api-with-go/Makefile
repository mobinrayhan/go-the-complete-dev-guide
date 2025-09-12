migrate-up:
	go run cmd/migrate/main.go -up
	
migrate-down:
	go run cmd/migrate/main.go -down 1

migrate-force:
	go run cmd/migrate/main.go -force 1

migrate-version:
	go run cmd/migrate/main.go -version

migrate-new:
	migrate create -ext sql -dir ./internals/db/migrations -seq $(name)
