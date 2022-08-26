package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	taskController "backend/controllers"
	docs "backend/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ======================================================================
// Main
// ======================================================================
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	docs.SwaggerInfo.BasePath = "api/v1"

	v1 := router.Group("/api/v1")

	{ // Tasks
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", taskController.Get)
			tasks.GET(":id", taskController.GetByID)
			tasks.POST("", taskController.Post)
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

// @host localhost:8080/
// @BasePath /api/v1
func main() {
	router := setupRouter()

	router.Run()
}
