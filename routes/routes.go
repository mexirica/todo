package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mexirica/todo/handlers"
	"github.com/mexirica/todo/infra"
	"github.com/mexirica/todo/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func initializeRoutes(router *gin.Engine) {
	infra.ConnectDatabase()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handlers.LoginHandler)
			auth.POST("/signup", handlers.SignUpHandler)
		}
		users := v1.Group("/users")
		users.Use(middlewares.AuthMiddleware())
		{
			users.GET("", handlers.UserByIDHandler)
		}
		tasks := v1.Group("/tasks")
		tasks.Use(middlewares.AuthMiddleware())
		{
			tasks.POST("/", handlers.CreateTaskHandler)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
