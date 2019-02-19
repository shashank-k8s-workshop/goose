# Build
FROM golang:1.11 as build

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .
RUN go build ./cmd/goose-api
RUN go build ./cmd/health-check

# Release
FROM gcr.io/distroless/base

ENV PORT=8083

COPY --from=build /go/src/app/goose-api /
COPY --from=build /go/src/app/health-check /

HEALTHCHECK --interval=10s --timeout=3s CMD ["/health-check"]

ENTRYPOINT ["/goose-api"]

