# Proyecto Talent Pitch API

Esta es una API desarrollada con Go, utilizando arquitectura hexagonal. La API gestiona las entidades `users`, `challenges`, y `videos`, y está desplegada en Render.

## Tecnologías Utilizadas

- **Go**: Lenguaje de programación utilizado para desarrollar la API.
- **Echo**: Framework para la creación de APIs REST en Go.
- **GORM**: ORM (Object-Relational Mapping) para gestionar la base de datos.
- **Viper**: Librería para la lectura de configuraciones desde el archivo `.env`.
- **OpenAI API**: API externa utilizada para generar descripciones y títulos para las entidades `challenges` y `videos`.

## Configuración y Despliegue

### Variables de Entorno

El proyecto requiere una variable de entorno para la API key de OpenAI llamada `OPENAI_API_KEY`, la cual es mandatoria para que la aplicación arranque correctamente.

```sh
OPENAI_API_KEY=tu_openai_api_key
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=dbname
```

Nota: Aunque los valores de configuración para la base de datos se leen desde el archivo .env, en este proyecto no se utilizan ya que la base de datos es en memoria. Esta configuración es parte de los requisitos de la prueba.

### Despliegue

La API está desplegada en Render y se puede acceder a través de la siguiente URL:

https://talent-pitch-api.onrender.com

Nota especial: Este es un servicio gratuito, por la tanto la velocidad de respuesta está limitada y una petición puede tardar para ser atendida.

### Ejecución Local

Para ejecutar la API localmente, asegúrate de tener Go instalado y sigue los siguientes pasos:

```
git clone https://github.com/javierBros/talent-pitch-api.git
cd talent-pitch-api
go run main.go
```

## Funcionamiento

Al arrancar la aplicación, se carga en una goroutine los 30 registros por cada tabla (users, challenges, videos). Se consume el API de OpenAI GPT para llenar los campos title y description de las tablas challenges y videos.

## Endpoints

### Users

* `POST /v1/users` Crear un nuevo usuario.

```
curl -X POST https://talent-pitch-api.onrender.com/v1/users -H "Content-Type: application/json" -d '{"name":"John Doe","email":"john.doe@example.com"}'
```
* `GET /v1/users/:id` Obtener un usuario por su ID.

```
curl https://talent-pitch-api.onrender.com/v1/users/1
```
* `GET /v1/users` Listar usuarios con paginación.

```
curl https://talent-pitch-api.onrender.com/v1/users?limit=10&offset=0
```
* `DELETE /v1/users/:id` Eliminar un usuario por su ID.

```
curl -X DELETE https://talent-pitch-api.onrender.com/v1/users/1
```
### Challenges

* `POST /v1/challenges` Crear un nuevo desafío.

```
curl -X POST https://talent-pitch-api.onrender.com/v1/challenges -H "Content-Type: application/json" -d '{"title":"Challenge 1","description":"Description 1","difficulty":3,"user_id":1}'
```
* `GET /v1/challenges/:id` Obtener un desafío por su ID.

```
curl https://talent-pitch-api.onrender.com/v1/challenges/1
```
* `GET /v1/challenges` Listar desafíos con paginación.

```
curl https://talent-pitch-api.onrender.com/v1/challenges?limit=10&offset=0
```
* `DELETE /v1/challenges/:id` Eliminar un desafío por su ID.

```
curl -X DELETE https://talent-pitch-api.onrender.com/v1/challenges/1
```
### Videos

* `POST /v1/videos` Crear un nuevo video.

```
curl -X POST https://talent-pitch-api.onrender.com/v1/videos -H "Content-Type: application/json" -d '{"title":"Video 1","description":"Description 1","url":"http://example.com/video1","user_id":1}'
```
* `GET /v1/videos/:id` Obtener un video por su ID.

```
curl https://talent-pitch-api.onrender.com/v1/videos/1
```
* `GET /v1/videos` Listar videos con paginación.

```
curl https://talent-pitch-api.onrender.com/v1/videos?limit=10&offset=0
```
* `DELETE /v1/videos/:id` Eliminar un video por su ID.

```
curl -X DELETE https://talent-pitch-api.onrender.com/v1/videos/1
```