package main

import (
	"net/http"
	"strconv"
	u "unsafe"

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

func getObstacle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, int(u.Sizeof(Obstacle{}.ID)))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid obstacle ID"})
		return
	}

	for _, obstacle := range obstacles {
		if obstacle.ID == id {
			c.IndentedJSON(http.StatusOK, obstacle)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "obstacle not found"})
}

func deleteObstacle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, int(u.Sizeof(Obstacle{}.ID)))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid obstacle ID"})
		return
	}

	for i, obstacle := range obstacles {
		if obstacle.ID == id {
			obstacles = append(obstacles[:i], obstacles[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "obstacle not found"})
}

func postObstacle(c *gin.Context) {
	var newObstacle Obstacle

	if err := c.BindJSON(&newObstacle); err != nil {
		return
	}

	obstacles = append(obstacles, newObstacle)
	c.IndentedJSON(http.StatusCreated, newObstacle)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/obstacles", getObstacles)
	router.GET("/obstacles/:id", getObstacle)
	router.POST("/obstacles", postObstacle)
	router.DELETE("/obstacles/:id", deleteObstacle)
	return router
}

func main() {
	router := setupRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
