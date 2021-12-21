FROM golang AS builder
WORKDIR /src
# Download dependencies
COPY go.mod go.sum /
RUN go mod download
# Add source code
COPY . .
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN go build -o main .

# Multi-Stage production build
FROM alpine AS production
RUN apk --no-cache add ca-certificates

WORKDIR /app
# Retrieve the binary from the previous stage
COPY --from=builder /src/main .
RUN chmod +x main
# Expose port
EXPOSE 8080
# Set the binary as the entrypoint of the container
CMD ["./main"]
