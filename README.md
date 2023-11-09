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

## Docker

### Docker-images

1. Postgres postgres:14-alpine
2. Auth service:
    - auth-dev
    - auth-prod
3. Migration

### Compose

**Prod**

Files:
- ./infrastructure/db-prod.yml
- ./infrastructure/srv-prod.yml

Network:
- auth-prod

Volumes:
- auth-postgres-prod

**Dev**

Files:
- ./infrastructure/db-dev.yml
- ./infrastructure/srv-dev.yml

Network:
- auth-dev

Volumes:
- auth-postgres-dev
