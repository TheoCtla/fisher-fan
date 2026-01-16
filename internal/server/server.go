package server

import (
	"fmt"

	routesV1 "fisherman/internal/api/v1/routes"
	routesV2 "fisherman/internal/api/v2/routes"
	"fisherman/internal/variables"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1")
	v2 := apiGroup.Group("/v2")
	{
		healthV1 := v1.Group("/health")
		boatV1 := v1.Group("/boats")
		routesV1.Health(healthV1)
		routesV1.RegisterBoatRoutes(boatV1)

		healthV2 := v2.Group("/health")
		routesV2.Health(healthV2)
	}

	err := router.Run(fmt.Sprintf("%s:%s", variables.Address, variables.Port))
	if err != nil {
		panic(err)
	}
}
