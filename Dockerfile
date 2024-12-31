# Base image untuk Go
FROM golang:1.22 as builder

# Set working directory
WORKDIR /app

# Copy semua file ke container
COPY . .

# Install dependencies
RUN go mod tidy

# Build binary untuk setiap service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/api-gateway /app/cmd/api-gateway/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/media /app/cmd/media/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/product /app/cmd/product/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/user /app/cmd/user/main.go

# RUN go build -o bin/order ./cmd/order/main.go
# RUN go build -o bin/payment ./cmd/payment/main.go
# RUN go build -o bin/shipping ./cmd/shipping/main.go

# Stage kedua: final image
FROM alpine:latest

# Set working directory
WORKDIR /app

RUN apk add --no-cache bash

# Copy binary dari stage builder
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/config /app/config
# COPY --from=builder /app/sql /app/sql
# COPY --from=builder /app/scripts /app/scripts

# Expose port (API Gateway default: 5000)
EXPOSE 5000

# Default
CMD ["/app/bin/api-gateway"]