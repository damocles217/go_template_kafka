package server

import (
	"github.com/damocles217/microservice/server/middlewares"
	v1 "github.com/damocles217/microservice/server/routes/v1"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	router := gin.New()

	router.Use(middlewares.CORSMiddleware())

	superGroup := router.Group("/api")
	v1Group := superGroup.Group("/v1")

	v1.SetupUserRoutes(v1Group)

	return router
}
