FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o graphql-stream  main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/graphql-stream  /app/graphql-stream
ENTRYPOINT ["/app/graphql-stream"]

EXPOSE 8080/tcp
