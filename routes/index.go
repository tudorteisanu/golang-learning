package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	bookRoutes(superRoute)
	userRoutes(superRoute)
}
