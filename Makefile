.DEFAULT_GOAL := build
.PHONY: goose-api

PORT ?= 8083

build-local: 
	rm -rf ./goose-api 
	go build ./cmd/goose-api

start-local: goose-api
	./goose-api

build:
	docker build -t k8s-goose .

start:
	@docker run -d -p $(PORT):$(PORT) -e PORT=$(PORT) --name k8s-goose k8s-goose
 
stop:
	docker stop k8s-goose
	docker rm k8s-goose

logs: 
	docker logs -f k8s-goose