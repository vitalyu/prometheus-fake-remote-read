# Prometheus fake remote read storage

Remote read storage uses [promtool series notation](https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/#series)

### Artifacts

* [Binary releases](https://github.com/vitalyu/prometheus-fake-remote-read/releases)
* [Docker images](https://github.com/vitalyu/prometheus-fake-remote-read/pkgs/container/prometheus-fake-remote-read)


### Demo

1. Clone repo and run. It starts docker compose with prometheus and fake remote-read util
    ```
    make demo
    ```
2. Open [local prometheus](http://127.0.0.1:9090/graph?g0.expr=test&g0.tab=0&g0.display_mode=lines&g0.show_exemplars=0&g0.range_input=1h) (http://127.0.0.1:9090) and checkout `test` metric


### Instalation

1. Prepare configuration (see [example.config.json](./configs/example.config.json)). Please read [promtool series notation](https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/#series)
    ```
    {
        "input_series": [
        {
            "interval": "1m",
            "series":   "test{job='backfiller_test'}",
            "values":   "0+1x100 99-1x99"
        },
        {
            "interval": "10m",
            "series":   "test{job='backfiller_test2'}",
            "values":   "0+5x100"
        }
        ]
    }
    ```

2. Run remote_read server 
    ```
    ./prometheus-fake-remote-read --config ./your_config.json
    ```


3. Add remote_read address to your prometheus configuration and restart prometheus

    ```
    ...

    remote_read:
    - url: "http://vscode:9999/read"

    ...

    ```

4. For now you can query `test` metric. To update `prometheus-fake-remote-read` configuration you need to restart util. No need to restart prometheus.


### Development

Run vscode with devcontainers. Prometheus and Grafana presents.