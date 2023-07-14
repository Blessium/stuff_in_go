FROM golang:alpine AS dep-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


FROM golang:alpine AS builder

WORKDIR /app
COPY --from=dep-builder /app ./
COPY ./ ./
RUN mkdir -p /out/bin
RUN go build -o ./out/bin/rest_api ./cmd/rest_api/main.go


FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/out/bin/rest_api ./
EXPOSE 3000
CMD ["./rest_api"]
