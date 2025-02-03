# Etapa de build
FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLE=0 GOOS=linux go build -ldflags="-s -w" -o server ./src/cmd

# Etapa de execução
FROM scratch
WORKDIR /app/
COPY --from=builder /app/server .
CMD ["./server"]
