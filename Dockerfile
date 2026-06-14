# Multi-stage build - Stage 1 compila, Stage 2 só tem o binário. Imagem final ~10-15MB
# =============================================================================
# Stage 1: BUILD
# =============================================================================
FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY app/go.mod ./

# Baixa as dependências (no caso do meu desafio não tem externas, mas deixei por ser uma boa prática)
RUN go mod download

COPY app/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# =============================================================================
# Stage 2: RUNTIME
# =============================================================================
FROM alpine:3.19

RUN adduser -D -s /bin/sh appuser

WORKDIR /app

COPY --from=builder /build/server .

USER appuser

EXPOSE 8080

CMD ["./server"]
