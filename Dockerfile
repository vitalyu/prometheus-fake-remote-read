FROM alpine:latest
WORKDIR /root/

COPY prometheus-fake-remote-read .
COPY configs/example.config.json .

RUN ls -lh

ENTRYPOINT ["./prometheus-fake-remote-read"]
CMD ["--config", "./example.config.json"]
