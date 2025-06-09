# Build Stage
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum* ./
RUN go mod download
COPY . .
RUN go build -o minimal-webapp-with-template .

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/minimal-webapp-with-template .
EXPOSE 8080
CMD ["./minimal-webapp-with-template"]
