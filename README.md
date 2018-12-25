# Pgpool2 Prometheus Exporter

## WIP

## Installation
* The project can be built with Docker:
```
docker build -t pgpool-prometheus-exporter:0.0.1-dev .
```

and after that the container can be started with:
```
docker run -e PCP_USER=<pcp_user> -e PCP_PASSWORD=<pcp_password> -d -P pgpool-prometheus-exporter:0.0.1-dev
```

### Supported environment variables
The Docker image contains pgpool's pcp-binaries that it will invoke and parse the output. You can use following environment variables to configure invocation of pcp binaries:
* PCP_LOCATION - Location of pcp-binaries
* PCP_USER - Username for PCP process
* PCP_PASSWORD - Password for PCP process
* PCP_HOST - Host where pcp process is running
* PCP_PORT - Port where pcp process is running

### Usage
Since the exporter requires running pgpool and pcp process, running it standalone doesn't make much sense. Instead exporter is designed to run as Kubernetes' sidecar container. Instructions how to do that will be added later.


## Exported metrics
* node_count
