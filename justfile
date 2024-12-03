create_migration SEQ_NAME:
	migrate create -ext sql -dir migrations -seq {{SEQ_NAME}}

migrate POSTGRESQL_URL: # example: "postgres://postgres_user:postgres_pass@localhost:5432/benchmarking?sslmode=disable"
	migrate -database {{POSTGRESQL_URL}} -path migrations up

run: # Debug
    ENV="debug" go run ./cmd/server/main.go
    
watch:
    ENV="debug" reflex -r '\.go' -s -- sh -c 'go run ./cmd/server/main.go'

release:
    go build -ldflags "-s -w" -o ./bin/server ./cmd/server/main.go