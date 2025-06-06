# API REST en Go con Fiber y MongoDB

Este proyecto es una API REST construida con el framework [Fiber](https://gofiber.io/) en Go y MongoDB como base de datos.

## Características

- Autenticación con JWT (tokens válidos por 10 minutos)
- CRUD de usuarios y tareas
- Validación de datos
- MongoDB como base de datos
- Middleware para proteger rutas

## Tecnologías utilizadas

- Go v1.24.3
- Fiber v2.52.8
- MongoDB v7.0.5
- JWT (github.com/golang-jwt/jwt/v5) v5.2.2

## Instalación

1. Clona el repositorio:
- git clone https://github.com/DoloresSalinas/api-rest.git
- cd api-rest

2. Instala las dependencias:
- go mod tidy

3. Asegúrate de tener una base de datos MongoDB corriendo (local o remota).
Si tienes instalado MongoDB Compass, conéctate a:
- mongodb://localhost:27017
Y busca la base de datos (puede ser api_rest o con el nombre asignado), colección (puede ser users o con el nombre asignado).

4. Configura la conexión en config/database.go.

5. Ejecuta el servidor (en la terminal):
- go run main.go

O puedes compilar sin ejecutar directamente (por si el antivirus detecta solo al correr en caso de ser local)
- go build -o app.exe
- ./app.exe

## Autenticación

Para obtener el token JWT se requere usar el endpoint de login, al obtenerlo se debe enviar en las rutas protegidas usando el header:
- Authorization: Bearer <tu_token>

## Endpoints principales

| Método | Ruta            | Descripción               |
| ------ | --------------- | ------------------------- |
| POST   | /api/register   | Crear una nuevo usuario   |
| POST   | /api/login      | Iniciar sesión            |
| PUT    | /api/users/\:id | Actualizar un usuario     |
| DELETE | /api/users/\:id | Eliminar un usuario       |
| GET    | /api/users/\:id | Obtener un usuario        |
| GET    | /api/users      | Obtener todos los usuarios|
| POST   | /api/tasks      | Crear una nueva tarea     |
| PUT    | /api/tasks/\:id | Actualizar una tarea      |
| DELETE | /api/tasks/\:id | Eliminar una tarea        |
| GET    | /api/tasks/\:id | Obtener una tarea         |
| GET    | /api/tasks      | Obtener todas las tareas  |

## Autor
- María Dolores Salinas Jiménez





