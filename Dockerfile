# Stage 1: Build the Go application
FROM golang:1.21-alpine as builder
LABEL authors="dmmoody"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code and build
COPY . ./
RUN CGO_ENABLED=0 go build -o /main -ldflags="-s -w" .

# Stage 2: Create a minimal production image
FROM alpine:latest as prod
COPY --from=builder /main /main
ENTRYPOINT ["/main"]
