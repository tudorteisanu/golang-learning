package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(superRoute *gin.Engine) {
	superRoute.Use(XssMiddleware())
	if gin.Mode() == gin.ReleaseMode {
		superRoute.Use(SecurityMiddleware())
	}
	superRoute.Use(CorsMiddleware())
}
