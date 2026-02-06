package routes

import (
	"fisherfan/internal/api/v1/handlers"
	"fisherfan/internal/api/v1/repository"
	"fisherfan/internal/api/v1/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupLogRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewLogRepository(db)
	service := services.NewLogService(repo)
	handler := handlers.NewLogHandler(service)

	router.GET("/log/:id", handler.GetLogByUserID)
	router.POST("/log/:id", handler.CreateLogByUserID)
	router.GET("/log/:id/pages/:page_id", handler.GetPage)
	router.PATCH("/log/:id/pages/:page_id", handler.UpdatePage)
	router.DELETE("/log/:id/pages/:page_id", handler.DeletePage)

}
