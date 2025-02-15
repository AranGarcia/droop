# Droop

Herramienta para jugar D&D

> :construction: Actualmente en construción. 

# Requisitos

- Docker
- Docker Compose
- Go 1.24.0
- [buf CLI v.1.44.0](https://buf.build/docs/ecosystem/cli-overview)

# Instalación

Compila los protobufs

```bash
make proto
```

# Ejecución

Se pueden levantar todos los servicios con Docker Compose:

```sh
docker-compose up
```