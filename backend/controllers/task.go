package taskController

import (
	"backend/db/data"
	"backend/db/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get godoc
// @Summary Get all tasks
// @Description get all tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Task
// @Router /tasks [get]
func Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Tasks)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	for _, t := range data.Tasks {
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

func Post(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	data.Tasks = append(data.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
