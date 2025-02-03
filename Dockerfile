FROM golang:1.23.0-alpine3.20 AS builder

# Install tzdata if westack time filters are needed
RUN apk add --no-cache ca-certificates upx git # tzdata

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/fredyk/westack-go/v2@v2.0.0-beta

COPY . .

RUN westack-go generate
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/bin/diagnose-cloud

# Compress the binary only if we need to save space. It consumes more RAM on startup
# RUN upx --best --lzma /usr/bin/diagnose-cloud

RUN chmod +x /usr/bin/diagnose-cloud

RUN mkdir -p data

FROM scratch

WORKDIR /usr/src/app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Remember to copy zoneinfo if westack time filters are needed
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /usr/bin/diagnose-cloud /usr/bin/diagnose-cloud
COPY --from=builder /usr/src/app/data data/

COPY server/adapters/westack/entities/*.json server/adapters/westack/entities/
COPY server/*.json server/

EXPOSE 8080

ENV GO_ENV=DOCKER

ENTRYPOINT ["/usr/bin/diagnose-cloud"]
