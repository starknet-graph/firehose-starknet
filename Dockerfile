FROM golang:1.22-bullseye AS build

ARG VERSION

COPY . /src
WORKDIR /src

RUN go build -ldflags "-X main.version=$VERSION" ./cmd/firestark

FROM debian:bullseye-slim

LABEL org.opencontainers.image.source=https://github.com/starknet-graph/firehose-starknet

COPY --from=build /src/firestark /usr/bin/firestark

ENTRYPOINT ["/usr/bin/firestark"]
