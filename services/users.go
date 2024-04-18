package services

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/infra"
	"github.com/mexirica/todo/models"
	"net/http"
)

func GetByEmail(c *gin.Context) *models.User {
	var user *models.User
	email := c.Query("email")
	var id uint
	infra.DB.Model(&models.User{}).Where("email = ?", &email).Select("id").First(&id)
	if err := infra.DB.Preload("Tasks").First(&user, &id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return nil
	}
	return user
}
