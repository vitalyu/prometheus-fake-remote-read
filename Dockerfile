FROM alpine:latest
WORKDIR /root/

COPY prometheus-fake-remote-read .
COPY configs/example.config.json .

ENTRYPOINT ["./prometheus-fake-remote-read"]
CMD ["--config", "./example.config.json"]
