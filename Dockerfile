# Stage 1: Build the Go binary
FROM golang:1.20 AS build
WORKDIR /app

# Copy go.mod and go.sum separately to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o go-webapp .

# Stage 2: Create a smaller image for runtime
FROM alpine:latest
WORKDIR /app

# Copy only the necessary artifacts from the build stage
COPY --from=build /app/go-webapp .

# Expose the port your application runs on
EXPOSE 8080

# Command to run your application
CMD ["./go-webapp"]
