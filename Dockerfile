FROM --platform=$BUILDPLATFORM golang:1.23.6-alpine AS builder
LABEL maintainer="Tran Cong <trancong12102@gmail.com>"

ENV CGO_ENABLED 0

WORKDIR /app

RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

RUN mkdir -p /emptydir
RUN apk add --no-cache git ca-certificates

COPY go.* ./

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache go mod download

FROM builder AS build-image

COPY . .

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown
ARG REPOSITORY=unknown

ARG TARGETOS
ARG TARGETARCH

ENV GOOS ${TARGETOS}
ENV GOARCH ${TARGETARCH}

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
    go build -o /app/server \
    -ldflags="-s -w -X github.com/${REPOSITORY}/cmd.Version=${VERSION} -X github.com/${REPOSITORY}/cmd.Commit=${COMMIT} -X github.com/${REPOSITORY}/cmd.Date=${DATE}" \
    ./cmd/server

FROM scratch AS release
LABEL maintainer="Tran Cong <trancong12102@gmail.com>"

WORKDIR /app

COPY --from=build-image /emptydir /tmp
COPY --from=build-image /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-image /app/ /app/

ENV APP_ENV production
ENTRYPOINT ["/app/server"]
