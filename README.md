# http-debug

http-debug has some small tools to help you test and debug proxies and validate routing.

## Installation

```shell
# go
go install github.com/stenic/http-debug@latest

# docker 
docker pull ghcr.io/stenic/http-debug:latest

# dockerfile
COPY --from=ghcr.io/stenic/http-debug:latest /http-debug /usr/local/bin/
```

> For even more options, check the [releases page](https://github.com/stenic/http-debug/releases).


## Run

```shell
# Installed
http-debug -h

# Docker
docker run -ti ghcr.io/stenic/http-debug:latest -h

# Kubernetes
kubectl run http-debug --image=ghcr.io/stenic/http-debug:latest --restart=Never -ti --rm -- -h
```

## Documentation

```shell
http-debug -h
```

## Badges

[![Release](https://img.shields.io/github/release/stenic/http-debug.svg?style=for-the-badge)](https://github.com/stenic/http-debug/releases/latest)
[![Software License](https://img.shields.io/github/license/stenic/http-debug?style=for-the-badge)](./LICENSE)
[![Build status](https://img.shields.io/github/workflow/status/stenic/http-debug/Release?style=for-the-badge)](https://github.com/stenic/http-debug/actions?workflow=build)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## License

[License](./LICENSE)
