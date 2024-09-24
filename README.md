# Droop

Herramienta para jugar D&D

> :construction: Actualmente en construción. 

# Requisitos

- Docker y Docker Compose
- Go 1.22.6
- Protocol Buffers
  - protoc
  - protoc-gen-go

# Ejecución

```sh
# Esto debe levantar la base de datos
docker-compose up

# Ejecutar el servidor HTTP
go run ./characters/exec/http
```

> :warning: El binario del servidor se implementará en un contenedor eventualmente para poder ejeuctarlo con Docker Compose en local sin necesidad de requerir Go en local.