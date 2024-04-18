package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/services"
)

func SignUpHandler(c *gin.Context) {
	services.SignUp(c)
}

func UserByIDHandler(c *gin.Context) {
	user := services.GetByEmail(c)
	if user == nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(200, user)
}
