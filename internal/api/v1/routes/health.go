package routes

import (
	"fisherfan/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func Health(routerGrp *gin.RouterGroup) {
	routerGrp.GET("/", handlers.Health)
}
