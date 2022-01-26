FROM golang:1.16.0-alpine as builder

WORKDIR /app

COPY go.* ./

RUN set -ex; \
    apk update; \
    apk add --no-cache \
    git \
    gcc \
    musl-dev \
    make

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goapp cmd/main.go

##########################################

FROM alpine:3.13.1 as release

ARG APP_ENV
ARG VAULT_TOKEN
ARG APP_VERSION

ENV APP_ENV=${APP_ENV} \
    VAULT_TOKEN=${VAULT_TOKEN} \
    APP_VERSION=${APP_VERSION}

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/goapp ./goapp
COPY --from=builder /app/configs ./configs

EXPOSE 8080

CMD ["./goapp"]