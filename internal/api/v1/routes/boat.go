package routes

import (
	"fisherman/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterBoatRoutes(rg *gin.RouterGroup) {
	boatGroup := rg.Group("/boats")
	{
		boatGroup.GET("/", handlers.GetBoatsHandler)
	}
}
