# ----------- Build Stage -----------
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git tzdata

# Copy go.mod and go.sum early for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build from the correct main.go location
RUN go build -o profile-service ./cmd/myiradat-backend-user

# ----------- Final Stage -----------
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache tzdata

# Copy the built binary
COPY --from=builder /app/profile-service .

EXPOSE 7792

CMD ["./profile-service"]
