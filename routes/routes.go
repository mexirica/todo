package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/handlers"
	"github.com/mexirica/todo/infra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	infra.ConnectDatabase()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.LoginHandler)
		}
		users := v1.Group("/users")
		{
			users.POST("/", handlers.SignUpHandler)
		}

	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
