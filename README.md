# diagnose-cloud

## Introducción

Este proyecto es la implementación básica del backend de una aplicación de diagnóstico de enfermedades.

## Requisitos

- Go 1.23
- Docker
- Docker Compose

## Instalación y ejecución

1. Clonar el repositorio
2. Ejecutar `docker compose up -d --build` en la raíz del proyecto

## Endpoints

El servicio diagnose-cloud arrancará en el puerto 8080. Los endpoints disponibles está disponibles en swagger-ui, en la ruta http://localhost:8080/swagger

## Testing

Para ejecutar los tests, ejecutar el siguiente comando, que ya cargará previamente las variables de entorno necesarias y posteriormente lanzará la herramienta `go test`

```bash
./scripts/run_tests.sh
```