FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o ws-echo  main.go

FROM alpine
WORKDIR /app/
COPY --from=builder /app/ws-echo  /app/ws-echo
ENTRYPOINT ["/app/ws-echo"]

EXPOSE 8080
