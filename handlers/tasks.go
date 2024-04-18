package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/services"
)

func CreateTaskHandler(c *gin.Context) {
	services.CreateTask(c)
}
