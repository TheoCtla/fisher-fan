package routes

import (
	"fisherfan/internal/api/v2/handlers"

	"github.com/gin-gonic/gin"
)

func Health(routerGrp *gin.RouterGroup) {
	routerGrp.GET("/", handlers.Health)
}
