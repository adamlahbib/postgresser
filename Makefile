PROJECT_NAME=postgresser
MODULE_NAME=postgresser

.DEFAULT_GOAL := build

.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/proto/*.proto

.PHONY: run
run:
	@go run cmd/server/main.go

.PHONY: runc
runc:
	@go run cmd/client/main.go

.PHONY: build
build:
	@go build ./cmd/server/main.go

.PHONY: buildc
buildc:
	@go build ./cmd/client/main.go

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v -coverprofile coverage.out ./...

.PHONY: coverage
coverage:
	@go tool cover -html=coverage.out

.PHONY: get
get:
	@go mod download