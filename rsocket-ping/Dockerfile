FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o rsocketping main.go

FROM alpine
WORKDIR /app/
COPY --from=builder /app/rsocketping /app/rsocketping
ENTRYPOINT ["/app/rsocketping"]

EXPOSE 8080
