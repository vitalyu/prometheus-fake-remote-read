ARG TARGETOS
ARG TARGETARCH

FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o prometheus-fake-remote-read ./cmd/prometheus-fake-remote-read

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/prometheus-fake-remote-read .
COPY --from=builder /app/configs/example.config.json .

ENTRYPOINT ["./prometheus-fake-remote-read"]
CMD ["--config", "./example.config.json"]
