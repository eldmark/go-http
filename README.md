# go-http

A RESTful HTTP API built with Go (no external frameworks) that manages One Piece characters. Data is persisted to a local JSON file and the server runs inside Docker.

---

## Project Structure

```
go-http/
├── data/
│   └── onepiece.json        # JSON file used as the database
├── handlers/
│   └── handler_character.go # HTTP handler methods
├── models/
│   └── item.go              # Character model struct
├── utils/
│   └── items.go             # WriteJSON helper
├── main.go                  # Entry point & route registration
├── Dockerfile
└── docker-compose.yml
```

---

## Requirements

- [Docker](https://www.docker.com/) and Docker Compose

---

## Running the Server

```bash
docker compose build --no-cache
docker compose up
```

The server will be available at `http://localhost:24229`.

---

## Character Model

| Field       | Type   | JSON key      | Required |
|-------------|--------|---------------|----------|
| ID          | int    | `id`          | auto     |
| Name        | string | `name`        | yes      |
| Devil Fruit | string | `devil_fruit` | no       |
| Fight Style | string | `fight_style` | yes      |
| Weapon      | string | `weapon`      | yes      |
| Speciality  | string | `speciality`  | yes      |

---

## API Endpoints

### Health Check

```
GET /api/ping
```

**Response**
```json
{ "message": "pong" }
```

---

### Get All Characters

```
GET /api/characters
```

Supports optional query parameters for filtering:

| Param         | Description                             |
|---------------|-----------------------------------------|
| `id`          | Filter by exact ID                      |
| `name`        | Filter by exact name (case-insensitive) |
| `devil_fruit` | Filter by devil fruit (partial match)   |
| `weapon`      | Filter by weapon (partial match)        |
| `speciality`  | Filter by speciality (partial match)    |

**Examples**
```
GET /api/characters
GET /api/characters?name=Zoro
GET /api/characters?devil_fruit=gomu
```

---

### Get Character by ID

```
GET /api/characters/{id}
```

**Response `200`**
```json
{
  "id": 1,
  "name": "Monkey D. Luffy",
  "devil_fruit": "Gomu Gomu no Mi",
  "fight_style": "Elastic close-quarters combat",
  "weapon": "Body",
  "speciality": "Gear transformations"
}
```

**Response `404`**
```json
{ "error": "Character not found" }
```

---

### Add Character

```
POST /api/characters
Content-Type: application/json
```

**Request Body**
```json
{
  "name": "Boa Hancock",
  "devil_fruit": "Mero Mero no Mi",
  "fight_style": "Kick-based martial arts",
  "weapon": "Body",
  "speciality": "Petrification via love"
}
```

**Response `201`** — returns the created character with its generated `id`.

---

### Update Character

```
PUT /api/characters/{id}
Content-Type: application/json
```

**Request Body** — same fields as POST.

**Response `200`** — returns the updated character.

---

### Delete Character

```
DELETE /api/characters/{id}
```

**Response `200`**
```json
{ "message": "Character deleted" }
```

---

## Error Responses

All errors follow the same format:

```json
{ "error": "<description>" }
```

| Status | Meaning                                  |
|--------|------------------------------------------|
| 400    | Invalid input or missing required fields |
| 404    | Character not found                      |
| 405    | Method not allowed                       |

### Evidence

### Create Character

![Create Character](./public/createOne.png)

### Get All Characters

![Get All Characters](public/getCharacters.png)

### Get Single Character

![Get Single Character](public/getOneCharacter.png)
