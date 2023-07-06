FROM golang:alpine AS dep-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


FROM golang:alpine AS builder

WORKDIR /app
COPY --from=dep-builder /app ./
COPY ./ ./
RUN mkdir -p /out/bin
RUN go build -o ./out/bin/sample_prom ./cmd/sample_prom/main.go


FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/out/bin/sample_prom ./
EXPOSE 3000
CMD ["./sample_prom"]
