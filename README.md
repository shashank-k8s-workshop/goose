# Goose

A minimal golang service for k8s workshop

## Build
### Run on local machine
- build `make build-local`
- start `make start-local`

### Run in docker
- build `make build`
- start `make start`
- stop `make stop`
- teardown `make remove`
- tail logs `make logs`

## ENV Variables
- `PORT`: Port on which http server should listen. Default: `8083`