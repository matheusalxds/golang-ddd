# Etapa de build
FROM golang:1.23.1 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o api main.go

# Etapa de execução
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api .
CMD ["./api"]
