# Dorm selection microservice system.

## Environment

- Node.js
- docker-compose

``` shell
yum install -y nodejs

sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose

```

## Quick Start

1. Edit your server url in `Makefile` `VUE_APP_BASE_URL`
2. Following commands

```shell
make build
make serve
```
