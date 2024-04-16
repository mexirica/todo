package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/services"
)

func SignUpHandler(c *gin.Context) {
	services.SignUp(c)
}
