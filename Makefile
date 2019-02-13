.DEFAULT_GOAL := build

.PHONY: goose-api

build: 
	rm -rf ./goose-api 
	go build ./cmd/goose-api

run: goose-api
	./goose-api
