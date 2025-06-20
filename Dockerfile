# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# App runs on port 8080
EXPOSE 8080

CMD ["./main"]
