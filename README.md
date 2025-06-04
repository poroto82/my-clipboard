# My Clipboard

Una aplicación de sincronización de portapapeles que permite compartir el contenido del portapapeles entre diferentes máquinas usando Redis como backend.

## Características

- Sincronización en tiempo real del portapapeles
- Soporte para Wayland (wl-copy/wl-paste)
- Backend Redis para la comunicación entre instancias
- Configuración flexible del servidor Redis

## Requisitos

- Go 1.16 o superior
- Redis
- Wayland (wl-copy/wl-paste)
- Docker y Docker Compose (opcional, para ejecutar Redis)

## Instalación

1. Clona el repositorio:
```bash
git clone https://github.com/tu-usuario/my-clipboard.git
cd my-clipboard
```

2. Instala las dependencias:
```bash
go mod download
```

3. Compila el proyecto:
```bash
go build
```

## Uso

### Ejecutar con Docker Compose

1. Inicia Redis usando Docker Compose:
```bash
docker-compose up -d
```

2. Ejecuta la aplicación:
```bash
./my-clipboard
```

### Ejecutar manualmente

1. Asegúrate de tener Redis ejecutándose en tu sistema
2. Ejecuta la aplicación especificando la dirección de Redis:
```bash
./my-clipboard [redis_host:port]
```

Por defecto, la aplicación intentará conectarse a `192.168.2.3:6379`.

## Configuración

La aplicación acepta los siguientes parámetros:

- Sin argumentos: Usa la dirección por defecto (192.168.2.3:6379)
- `[redis_host:port]`: Especifica una dirección personalizada de Redis
- `--help` o `-h`: Muestra la ayuda

## Contribuir

Las contribuciones son bienvenidas. Por favor, abre un issue para discutir los cambios propuestos.

## Licencia

MIT 