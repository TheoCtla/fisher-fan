package server

import (
	"fmt"

	routesV1 "fisherfan/internal/api/v1/routes"
	routesV2 "fisherfan/internal/api/v2/routes"
	"fisherfan/internal/variables"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitServer(db *gorm.DB) {
	router := gin.Default()
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1")
	v2 := apiGroup.Group("/v2")
	{
		// Routes d'authentification (publiques et protégées)
		routesV1.SetupAuthRoutes(v1, db)

		// Routes utilisateurs (protégées par authentification)
		routesV1.SetupUserRoutes(v1, db)

		// Autres routes
		routesV1.SetupBoatRoutes(v1, db)
		routesV1.SetupTripRoutes(v1, db)
		routesV1.SetupReservationRoutes(v1, db)
		routesV1.SetupLogRoutes(v1, db)

		// Health checks
		routesV1.Health(v1)
		routesV2.Health(v2)
	}

	err := router.Run(fmt.Sprintf("%s:%s", variables.GlobalConfig.ServerAddress, variables.GlobalConfig.ServerPort))
	if err != nil {
		panic(err)
	}
}
