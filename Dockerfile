FROM golang:1.22.0 AS deps-installer

RUN apt-get update && apt-get install -y freetds-dev

WORKDIR /go/src/github.com/Dhouib-Mohamed/quoxy

COPY go.mod go.sum ./
RUN go mod download

COPY internal internal
COPY util util

FROM debian:bookworm-slim AS runner

RUN apt-get update && apt-get install -y freetds-bin

WORKDIR /app

COPY scripts/sql scripts/sql

#### Proxy builder
FROM deps-installer AS proxy-builder

COPY cmd/proxy cmd/proxy

RUN go build -o /go/bin/proxy ./cmd/proxy

FROM runner AS proxy

COPY --from=proxy-builder /go/bin/proxy .

CMD ["./proxy"]

#### REST API builder
FROM deps-installer AS rest-api-builder

COPY api/handler api/handler
COPY cmd/rest-api cmd/rest-api

RUN go build -o /go/bin/rest-api ./cmd/rest-api

FROM runner AS rest-api

COPY --from=rest-api-builder /go/bin/rest-api .

CMD ["./rest-api"]

#### Token Handler builder
FROM deps-installer AS token-handler-builder

COPY cmd/token-handler cmd/token-handler

RUN go build -o /go/bin/token-handler ./cmd/token-handler

FROM runner AS token-handler

COPY --from=token-handler-builder /go/bin/token-handler .

CMD ["./token-handler"]



