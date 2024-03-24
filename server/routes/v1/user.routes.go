package v1

import (
	"github.com/damocles217/microservice/server/controllers"
	"github.com/damocles217/microservice/server/database"
	"github.com/damocles217/microservice/server/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	userDb := database.InitDB()

	services := services.NewServicesWithDb(userDb)
	handler := controllers.SetNewHandler(services)

	group := r.Group("/user")
	{
		group.GET("/me", handler.GetUser)
	}
}
