FROM golang:1.22.4-bookworm as builder
# Create and change to the app directory.
WORKDIR /app

RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache go build -v -o server ./cmd/server

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server .

# Run the web service on container startup.
ENV APP_ENV=production
CMD ["/app/server"]
