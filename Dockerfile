# syntax=docker/dockerfile:1

# build
FROM golang:1.19 AS builder

WORKDIR /app

# copy go mod files
COPY go.mod go.sum ./

# download any dependencies
RUN go mod download

# copy source code
COPY . .

# build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

# copy the binary from builder
COPY --from=builder /app/main .

# expose port
EXPOSE 8080

# run the application
CMD ["./main"]