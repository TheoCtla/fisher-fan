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
		healthV2 := v2.Group("/health")

		routesV1.Health(healthV1)
		routesV2.Health(healthV2)
	}

	err := router.Run(fmt.Sprintf("%s:%s", variables.Adress, variables.Port))
	if err != nil {
		panic(err)
	}
}
