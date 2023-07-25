package routes

import (
	"github.com/gin-gonic/gin"
	"golang-rest-api-template/controllers"
)

func userRoutes(superRoute *gin.RouterGroup) {
	booksRouter := superRoute.Group("/users")
	{
		booksRouter.GET("/", controllers.FindUsers)
		booksRouter.POST("/", controllers.CreateUser)
		booksRouter.GET("/:id", controllers.FindUser)
		booksRouter.PUT("/:id", controllers.UpdateUser)
		booksRouter.DELETE("/:id", controllers.DeleteUser)
	}
}
