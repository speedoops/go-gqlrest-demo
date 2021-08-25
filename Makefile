.PHONY: gen run test all

all: gen build test lint

gql:
	@# https://github.com/golang/go/issues/44129 // fix go 1.16 bug
	go run -mod=mod github.com/99designs/gqlgen

rest:
	go run -mod=mod github.com/speedoops/gql2rest

gen: gql rest

build:
	go build

run:	
	# @command -v air &> /dev/null || go install github.com/cosmtrek/air
	# air
	go run server.go

test:
	# go test ./...
	go.exe test -timeout 30s -run ^TestTodo ./... -v

lint:
	golangci-lint run -v --timeout=5m
