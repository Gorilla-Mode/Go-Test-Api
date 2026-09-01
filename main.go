package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Obstacle struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Geometry string `json:"geometry"`
}

var obstacles = []Obstacle{
	{ID: 1, Name: "Obstacle 1", Geometry: "Polygon((0 0, 10 0, 10 10, 0 10))"},
	{ID: 2, Name: "Obstacle 2", Geometry: "Polygon((10 0, 20 0, 20 10, 10 10))"},
}

func getObstacles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, obstacles)
}

func main() {
	router := gin.Default()
	router.GET("/obstacles", getObstacles)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
