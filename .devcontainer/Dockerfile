FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# must be amd64 or arm64
ARG TARGETARCH

## promtool from prometheus
## https://github.com/prometheus/prometheus/releases
RUN cd /tmp/ &&
    prometheus_version=2.53.1 &&
    curl -L https://github.com/prometheus/prometheus/releases/download/v${prometheus_version}/prometheus-${prometheus_version}.linux-${TARGETARCH}.tar.gz | tar zx &&
    cp -f prometheus-${prometheus_version}.linux-${TARGETARCH}/promtool /usr/bin/ &&
    rm -rf prometheus-${prometheus_version}.linux-${TARGETARCH}
