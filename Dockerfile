FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.19

RUN adduser -D -s /bin/sh appuser

WORKDIR /app

COPY --from=builder /build/server .

USER appuser

EXPOSE 8080

CMD ["./server"]
