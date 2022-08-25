package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ======================================================================
// In-memory data
// ======================================================================

type Task struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Points float64 `json:"points"`
}

// Maybe?
type Action struct {
	ID     string  `json:"id"`
	TaskID string  `json:"taskId"`
	Points float64 `json:"points"`
}

// albums slice to seed record album data.
var tasks = []Task{
	{ID: "1", Title: "Wake up on time", Points: 10},
	{ID: "2", Title: "Study", Points: 25},
	{ID: "3", Title: "Exercise", Points: 25},
}

// ======================================================================
// Endpoints
// ======================================================================

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range tasks {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// ======================================================================
// Main
// ======================================================================

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.POST("/tasks", postTask)

	return router
}

func main() {
	router := setupRouter()

	router.Run()
}
