package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/services"
)

func LoginHandler(c *gin.Context) {
	services.Login(c)
}
