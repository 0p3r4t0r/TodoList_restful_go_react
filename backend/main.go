package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/db/models"
)

// ======================================================================
// In-memory data
// ======================================================================

// Maybe?
// albums slice to seed record album data.
var tasks = []models.Task{
	{Title: "Wake up on time", Points: 10},
	{Title: "Study", Points: 25},
	{Title: "Exercise", Points: 25},
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
		id_u64, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			fmt.Println(err)
		}

		if t.ID == uint(id_u64) {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postTask(c *gin.Context) {
	var newTask models.Task

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
