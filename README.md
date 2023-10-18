[![Run Tests](https://github.com/mixdjoker/auth/actions/workflows/auth.yaml/badge.svg)](https://github.com/mixdjoker/auth/actions/workflows/auth.yaml)
[![Deploy](https://github.com/mixdjoker/auth/actions/workflows/deploy.yaml/badge.svg)](https://github.com/mixdjoker/auth/actions/workflows/deploy.yaml)

# auth

## Installation

### Запуск локального Docker-образ

```shall
source .credentials
docker compose up -d
```

### Ручное копирование на удаленный сервер

```shell
sudo sed -i 's/^\(.*\)\scourse$/<<new_ip_address>> course/' /etc/hosts
make build
make copy-to-server
```

*new_ip_address* - IP Address сервера
