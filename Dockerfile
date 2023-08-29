FROM golang:1.20 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build server
RUN CGO_ENABLED=0 GOOS=linux go build -v -o gambling_platform_server ./cmd/backend/main.go

# Build client
RUN CGO_ENABLED=0 GOOS=linux go build -v -o gambling_platform_client ./cmd/client/main.go

FROM alpine:3.18

# Copy both server and client
COPY --from=builder /app/gambling_platform_server /gambling_platform_server
COPY --from=builder /app/gambling_platform_client /gambling_platform_client

# Define a default command (here, running the server)
CMD ["/gambling_platform_server"]
