package routes

import (
	"go-lang-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

// receives a router of type gin.Engine and returns a gin.Engine
func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1") // add a prefix to the routes

	{
		books := main.Group("books") // /api/v1/books
		{
			books.GET("/", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router

}
