FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o graphql-rest main.go

FROM alpine
WORKDIR /app/
COPY --from=builder /app/graphql-rest /app/graphql-rest
ENTRYPOINT ["/app/graphql-rest"]

EXPOSE 8080

