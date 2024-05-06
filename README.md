# Glamping API Project 🚀

## Descripción 📘
Este proyecto es una API en Go que maneja la autenticación y registro de usuarios. Utiliza Gin para el manejo de rutas HTTP y MySQL como sistema de gestión de bases de datos. La API soporta operaciones para el registro y login de usuarios y algunas otras operaciones.

## Características 🌟
- **Login de Usuarios**: Permite a los usuarios iniciar sesión.
- **Registro de Usuarios**: Permite a nuevos usuarios registrarse.
- **Manejo de CORS**: Configuración de CORS para aceptar solicitudes de dominios específicos.

## Comenzando 🏁

### Pre-requisitos 📋
- Go 1.15 o superior
- MySQL 5.7 o superior
- Git

### Instalación 🔧

1. Clona el repositorio:
   ```bash
   git clone https://github.com/iTzKevinPG/glamping-api.git
   ```

2. Navega al directorio del proyecto:
   ```bash
   cd glamping-api
   ```

3. Instala las dependencias:
   ```bash
   go mod tidy
   ```

## Configuración de la base de datos 🗄️

1. Crea una base de datos MySQL llamada glampingdb.
2. Ejecuta el script SQL que se encuentra en ./migrations/ para configurar la tabla necesaria o el api automáticamente las crea.

## Ejecución 🚀

Para iniciar el servidor, ejecuta:

```bash
go run cmd/main.go
   ```

El servidor se iniciará en localhost:8080 y estará listo para aceptar solicitudes.

## Dependencias 🛠️
- Gin: Framework web utilizado para manejar las solicitudes HTTP.
- MySQL Driver: Driver de MySQL para Go, usado para conectar y operar con la base de datos.

## 🎁 Expresiones de Gratitud

- Comenta a otros sobre este proyecto 📢
- Invita una cerveza 🍺 o un café ☕ a alguien del equipo.
