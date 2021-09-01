.PHONY: all gen build run test lint

all: gen build test lint

gen:
	@# https://github.com/golang/go/issues/44129 // add "-mod=mod" as workaround for go 1.16 bug
	@rm graph/generated/rest.go &>/dev/null || true
	go run -mod=mod github.com/99designs/gqlgen
	go run -mod=mod github.com/speedoops/go-gqlrest

build:
	go build

run:	
	# @command -v air &>/dev/null || go install github.com/cosmtrek/air
	# air
	go run server.go

test:
	# go test ./...
	go.exe test -timeout 30s -run ^TestTodo ./... -v

lint:
	golangci-lint run -v --timeout=5m
