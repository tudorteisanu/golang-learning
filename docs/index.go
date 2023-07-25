package docs

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupSwagger(superRoute *gin.RouterGroup) {

	SwaggerInfo.BasePath = "/api/v1"
	superRoute.GET("/api/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
