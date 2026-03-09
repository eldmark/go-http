# 09 - Query Parameters

En esta etapa la API incorpora soporte para parámetros en la URL, permitiendo filtrar resultados utilizando query parameters.

Se mantiene el modelo basado en archivo JSON como fuente de datos, pero ahora el comportamiento del endpoint cambia según los parámetros recibidos.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Qué son los query parameters
* Cómo acceder a ellos con `r.URL.Query()`
* Cómo convertir valores de string a otros tipos (`strconv.Atoi`)
* Cómo validar input del usuario
* Cómo devolver diferentes respuestas según parámetros recibidos

---

## 📁 Estructura del proyecto

```
.
├── main.go
├── data/
│   └── teams.json
├── Dockerfile
└── docker-compose.yml
```

La estructura no cambia respecto a la rama anterior.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* `GET /api/teams` devolvía siempre todos los equipos

Ahora:

* `GET /api/teams` devuelve todos los equipos
* `GET /api/teams?id=1` devuelve un equipo específico

El mismo endpoint ahora tiene comportamiento condicional basado en parámetros.

---

## 🧩 Ejemplos de uso

Obtener todos los equipos:

```
GET /api/teams
```

Obtener un equipo específico:

```
GET /api/teams?id=1
```

Si el parámetro es inválido:

* Se devuelve `400 Bad Request`

Si el equipo no existe:

* Se devuelve `404 Not Found`

---

## 🔎 Conceptos introducidos

* `r.URL.Query()` para leer parámetros de la URL
* Uso de `query.Get("id")`
* Conversión de tipos con `strconv.Atoi`
* Validación de parámetros
* Respuestas condicionales según input

---

## 🐳 Ejecución

El servidor escucha en el puerto 80 dentro del contenedor.

En `docker-compose.yml` se mapea:

```yaml
ports:
  - "8080:80"
```

Probar con:

```bash
curl http://localhost:8080/api/ping
curl http://localhost:8080/api/teams
curl http://localhost:8080/api/teams?id=1
```

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa entendemos que:

* Los endpoints pueden cambiar su comportamiento según parámetros
* Todos los parámetros llegan como strings
* Es responsabilidad del backend validar y convertir los datos

Este es el paso previo antes de modelar recursos utilizando path parameters y construir una API REST más formal.

