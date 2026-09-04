# NRLCANICUS API

Test api in go using gin

## Curl commands

The API listens on `http://localhost:8080` when the application is running.

### 1. Get all obstacles

```shell
curl http://localhost:8080/obstacles \
    --header "Content-Type: application/json" \
    --request "GET"
```

### 2. Get an obstacle by ID

Replace `1` with the ID of the obstacle to retrieve.

```shell
curl http://localhost:8080/obstacles/1 \
    --header "Content-Type: application/json" \
    --request "GET"
```

### 3. Create an obstacle

The request body must contain an `id`, `name`, and `geometry`.

```shell
curl http://localhost:8080/obstacles \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": 3, "name": "Obstacle 3", "geometry": "Polygon((20 0, 30 0, 30 10, 20 10))"}'
```

### 4. Delete an obstacle

Replace `1` with the ID of the obstacle to delete.

```shell
curl http://localhost:8080/obstacles/1 \
    --include \
    --request "DELETE"
```

## PowerShell curl commands

### 1. Get all obstacles

```powershell
curl.exe http://localhost:8080/obstacles `
--header "Content-Type: application/json" `
--request "GET"
```

### 2. Get an obstacle by ID

```powershell
curl.exe http://localhost:8080/obstacles/1 `
--header "Content-Type: application/json" --request "GET"
```

### 3. Create an obstacle

```powershell
curl.exe http://localhost:8080/obstacles `
--include `
--header "Content-Type: application/json" `
--request "POST" `
--data '{"id": 3, "name": "Obstacle 3", "geometry": "Polygon((20 0, 30 0, 30 10, 20 10))"}'
```

### 4. Delete an obstacle

```powershell
curl.exe http://localhost:8080/obstacles/1 `
--include `
--request "DELETE"
```

## Possible Stack
| Type            | Name      | Description                                                                                  |
|-----------------|-----------|----------------------------------------------------------------------------------------------|
| Language        | Go        | Fast and simple statically typed compiled garbage collected language                         |
| Web             | Gin       | Fast HTTP web framework written in Go.                                                       |
| Auth**          | Hydra     | OAuth2/OIDC server that can run as a separate service and plug into any Go app               | 
| Auth**          | Authboss  | complete user system (signup, login, password reset, email/OAuth integration) in one library |
| Database        | Postgres  | Just use postgres.                                                                           |
| Database driver | Pgx       | Postgres driver for go                                                                       |  
| ORM***          | Sqlc      | Go Relational Mapping library                                                                |
| Migrations***   | Goose     | Database migrations library                                                                  |
| Audit db        | Cassandra | Open source, distributed NoSQL database. Nice for append only operations                     |
| Container       | Docker    | Containerization platform                                                                    |

**Assumes that self-hosted auth is preferable.
***Assumes a database-first approach.
