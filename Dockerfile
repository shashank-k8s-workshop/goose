# Build
FROM golang:1.11 as build

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .
RUN go build ./cmd/goose-api

# Release
FROM gcr.io/distroless/base
COPY --from=build /go/src/app/goose-api /
CMD ["/goose-api"]