FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o grpc-ping main.go && \
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.cert -subj "/CN=localhost/O=localhost"

FROM alpine:latest
WORKDIR /app/

COPY --from=builder /app/tls.cert /app/tls.cert
COPY --from=builder /app/tls.key /app/tls.key

COPY --from=builder /app/grpc-ping  /app/grpc-ping
ENTRYPOINT ["/app/grpc-ping"]

EXPOSE 8080/tcp
