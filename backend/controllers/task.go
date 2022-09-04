package taskController

import (
	"backend/db"
	"backend/db/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get godoc
// @Summary Get all tasks
// @Tags		tasks
// @Accept  	json
// @Produce  	json
// @Success 	200 {array}		models.Task
// @Router 		/tasks [get]
func Get(c *gin.Context) {
	var tasks []models.Task
	database, _ := db.Connect()
	database.Find(&tasks)
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetByID godoc
// @Summary		Get task by ID
// @Tags		tasks
// @Param  		id 	path 		int true "Task ID"
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} 	models.Task
// @Router 		/tasks/{id} 	[get]
func GetByID(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	database, _ := db.Connect()
	result := database.First(&task, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// Post godoc
// @Summary 	Create a new task
// @Tags		tasks
// @Param  		task	body models.CreateTaskDTO true "New Task"
// @Accept  	json
// @Produce  	json
// @Success 	201 {object} models.Task
// @Router 		/tasks [post]
func Post(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		return
	}

	database, _ := db.Connect()
	database.Select("Title", "Points").Create(&task)

	c.IndentedJSON(http.StatusCreated, task)
}
