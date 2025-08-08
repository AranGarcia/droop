# Droop

Herramienta para jugar D&D

> :construction: Actualmente en construción. 

## Requisitos

- Docker
- Docker Compose
- Go
- [buf CLI v.1.44.0](https://buf.build/docs/ecosystem/cli-overview)

## Instalación

Compila los protobufs

```bash
make proto
```

## Ejecución

Se pueden levantar todos los servicios con Docker Compose:

```sh
docker compose build
docker compose up
```

## Web

Para ejecutar la página web en local, revisa el directorio [`web/`](web/README.md)