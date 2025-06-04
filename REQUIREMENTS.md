# Requisitos del Sistema

Este documento detalla las dependencias del sistema necesarias para ejecutar My Clipboard.

## Dependencias Principales

### wl-clipboard
- **Descripción**: Herramientas de línea de comandos para interactuar con el portapapeles de Wayland
- **Comandos requeridos**: `wl-copy`, `wl-paste`
- **Versión mínima**: 2.0.0
- **Instalación**:
  - Fedora/RHEL: `sudo dnf install wl-clipboard`
  - Ubuntu/Debian: `sudo apt install wl-clipboard`
  - Arch Linux: `sudo pacman -S wl-clipboard`
  - Gentoo: `sudo emerge --ask wl-clipboard`

### Redis
- **Descripción**: Base de datos en memoria usada como backend para la sincronización
- **Versión mínima**: 6.0.0
- **Instalación**:
  - Fedora/RHEL: `sudo dnf install redis`
  - Ubuntu/Debian: `sudo apt install redis-server`
  - Arch Linux: `sudo pacman -S redis`
  - Gentoo: `sudo emerge --ask redis`

### Go
- **Descripción**: Lenguaje de programación
- **Versión mínima**: 1.16
- **Instalación**: Ver [instrucciones oficiales de Go](https://golang.org/doc/install)

## Verificación de Instalación

Para verificar que todas las dependencias están instaladas correctamente:

```bash
# Verificar wl-clipboard
wl-copy --version
wl-paste --version

# Verificar Redis
redis-cli --version

# Verificar Go
go version
```

## Notas Importantes

1. La aplicación requiere un entorno Wayland para funcionar correctamente
2. Si estás usando X11, necesitarás una alternativa como `xclip` o `xsel`
3. Redis puede ejecutarse localmente o en un servidor remoto
4. Para desarrollo, se recomienda usar Docker Compose para Redis 