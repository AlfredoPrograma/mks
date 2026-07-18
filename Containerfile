# ------ BUILDER STAGE -----
FROM golang:1.26.3 AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" -o mks .

# ------ MIGRATOR STAGE -----
FROM kukymbr/goose-docker:3.27.2 AS migrator

WORKDIR /migrations
COPY database/migrations .

# ------ RUNNER STAGE -----
FROM alpine:3.23.5 AS runner

WORKDIR /opt/mks

# Create runner user
RUN adduser -D -H mks
USER mks

# Bring built binary
COPY --from=builder --chown=mks:mks /app/mks .

# Run application
EXPOSE 8080
CMD ["./mks"]
