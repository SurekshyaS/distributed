# builder
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /app/netapp-storage-service ./cmd/server

# final
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/netapp-storage-service .
EXPOSE 8080
CMD ["./netapp-storage-service"]
