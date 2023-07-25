package routes

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/controllers"
)

func bookRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/books")
	{
		booksRouter.GET("/", controllers.FindBooks)
		booksRouter.POST("/", controllers.CreateBook)
		booksRouter.GET("/:id", controllers.FindBook)
		booksRouter.PUT("/:id", controllers.UpdateBook)
		booksRouter.DELETE("/:id", controllers.DeleteBook)
	}
}
