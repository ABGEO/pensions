FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN GOOS=linux go build -v ./cmd/pensions

FROM alpine:edge

LABEL org.opencontainers.image.title="Pensions"
LABEL org.opencontainers.image.authors="Temuri Takalandze <me@abgeo.dev>"
LABEL org.opencontainers.image.source="https://github.com/ABGEO/pensions"

WORKDIR /app

RUN addgroup --system --gid 1001 pensions \
    && adduser --system --uid 1001 pensions

COPY --from=builder --chown=pensions:pensions /app/pensions .
# @todo: remove this once Viper is fixed.
COPY --chown=pensions:pensions configs/example.env .env

USER pensions

ENTRYPOINT ["/app/pensions"]
