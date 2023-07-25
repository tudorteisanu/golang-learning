package main

import (
	"golang-rest-api-template/docs"
	"golang-rest-api-template/middleware"
	"golang-rest-api-template/models"
	"golang-rest-api-template/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	models.ConnectDatabase()
	middleware.SetupMiddleware(r)
	router := r.Group("/api/v1")
	routes.AddRoutes(router)
	docs.SetupSwagger(router)

	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
