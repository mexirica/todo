package services

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/infra"
	"github.com/mexirica/todo/models"
	"net/http"
)

func CreateTask(c *gin.Context) {
	id, _ := c.Get("id")
	if id == nil {
		c.JSON(500, gin.H{
			"message": "Error to read the user id",
		})
		return
	}
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON"})
		return
	}

	if description, ok := jsonData["description"].(string); ok {
		var task models.Task
		task.Description = description
		task.UserID = id.(float64)

		if err := infra.DB.Model(&models.Task{}).Create(&task).Error; err != nil {
			c.JSON(500, gin.H{
				"message": "Error to create the task",
			})
			return
		}
		c.JSON(201, gin.H{
			"message": "Task created successfully",
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campo 'description' não encontrado ou inválido"})
	}

}
