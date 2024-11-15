FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o udp-server  main.go

FROM alpine
WORKDIR /app/
COPY --from=builder /app/udp-server  /app/udp-server
ENTRYPOINT ["/app/udp-server"]

EXPOSE 8080/udp
