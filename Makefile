migrate_create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate:
	migrate -database "postgres://user:password@localhost:5433/tasks?sslmode=disable" -path ./migrations up

build:
	go build -o tasks .
