include: .env

local:
	go run cmd/main.go --env=local
start:
	go run cmd/main.go
build:
	go build cmd/main.go