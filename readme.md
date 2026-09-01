# NRLCANICUS API

Test api in go using gin

## Curl commands

1. GET 
```shell
curl http://localhost:8080/obstacles \
    --header "Content-Type: application/json" \
    --request "GET"
```

2. POST
```shell
curl http://localhost:8080/obstacles \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 3, "Name": "Obstacle 3", "Geometry": "Polygon((10 0, 20 0, 20 10, 10 10))"}'
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

**Assumes that self-hosted auth is preferable. Mutally exclusive.

***Assumes a database-first approach.
