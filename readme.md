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