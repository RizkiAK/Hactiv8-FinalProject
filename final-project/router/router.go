package router

import (
	"final-project/controller"
	"final-project/docs"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	// swagger embed files
)

func StartRouter() {
	route := gin.Default()

	docs.SwaggerInfo.Title = "Swagger TodoList"
	docs.SwaggerInfo.Description = "Documentation of TodoList Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost" + ":8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	todosRoute := route.Group("/todos")
	{
		todosRoute.POST("/", controller.TodoCreate)
		todosRoute.PUT("/:id", controller.TodoUpdate)
		todosRoute.DELETE("/:id", controller.TodoDelete)
		todosRoute.GET("/", controller.TodoGetAll)
		todosRoute.GET("/:id", controller.TodoGetById)
	}

	route.Run(":8080")
}
