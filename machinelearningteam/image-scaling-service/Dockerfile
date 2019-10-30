FROM golang:1.13 AS builder

WORKDIR /app


RUN wget -O grpc-health-probe -q https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.1/grpc_health_probe-linux-amd64 \
    && echo -n "e2ef834c3e8bf1efc1b8281532593cc647a246a042b7ae77b4a82ef9e6c70d22  grpc-health-probe" | sha256sum -c - \
    && chmod +x grpc-health-probe \
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the application
RUN go build cmd/imagescaler/imagescaler.go

FROM debian:stretch-slim AS runner

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

COPY --from=builder /app/imagescaler /app/imagescaler
COPY --from=builder /app/grpc-health-probe /usr/local/bin/grpc-health-probe

ENTRYPOINT ["/app/imagescaler"]
