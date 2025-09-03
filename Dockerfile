# Step 1: Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Copy ONLY go.mod and go.sum first (ensures caching works)
COPY go.mod go.sum ./

RUN go mod download

# Now copy the rest of the code
COPY . .

RUN go build -o server ./main.go

# Step 2: Runtime stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
