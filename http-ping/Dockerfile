FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o http-ping  main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/http-ping  /app/http-ping
ENTRYPOINT ["/app/http-ping"]

EXPOSE 8080/tcp
