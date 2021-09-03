.PHONY: all gen build run test lint

all: gen build test lint

gen:
	@# https://github.com/golang/go/issues/44129 // add "-mod=mod" as workaround for go 1.16 bug
	go run -mod=mod github.com/speedoops/go-gqlrest

build:
	go build

run:	
	# @command -v air &>/dev/null || go install github.com/cosmtrek/air
	# air
	go run main.go

test:
	# go test ./... -v
	go.exe test -timeout 30s -run ^TestTodo ./... -v

lint:
	golangci-lint run --timeout=5m
