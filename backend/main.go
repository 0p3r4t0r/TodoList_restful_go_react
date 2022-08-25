package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/db/models"

	docs "backend/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// getTasks godoc
// @Summary Get all tasks
// @Description get all tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Task
// @Router /tasks [get]
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
	docs.SwaggerInfo.BasePath = "api/v1"

	v1 := router.Group("/api/v1")

	{ // Tasks
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", getTasks)
			tasks.GET(":id", getTaskByID)
			tasks.POST("", postTask)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

// @title ToDo List API
// @version 1.0
// @description This is a sample project to learn API development in Golang.
// @termsOfService http://swagger.io/terms/

// @contact.name 0p3r4t0r
// @contact.url https://github.com/0p3r4t0r

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := setupRouter()

	router.Run()
}
