# ----------- Build Stage -----------
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Install git and tzdata
RUN apk add --no-cache git tzdata

# Copy go mod files and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o profile-service .

# ----------- Final Stage -----------
FROM alpine:latest

WORKDIR /app

# Install timezone support
RUN apk add --no-cache tzdata

# Copy the binary
COPY --from=builder /app/profile-service .

EXPOSE 7792

CMD ["./profile-service"]
