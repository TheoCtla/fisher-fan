package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterBoatRoutes(rg *gin.RouterGroup) {
	boatGroup := rg.Group("/boats")
	{
		boatGroup.GET("/", getBoatsHandler)
		boatGroup.POST("/", createBoatHandler)
		boatGroup.GET("/:id", getBoatByIDHandler)
		boatGroup.PUT("/:id", updateBoatHandler)
		boatGroup.DELETE("/:id", deleteBoatHandler)
	}
}
