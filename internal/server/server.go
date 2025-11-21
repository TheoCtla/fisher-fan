package server

import (
	"fmt"

	"fisherman/internal/variables"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	err := router.Run(fmt.Sprintf("%s:%s", variables.Adress, variables.Port))
	if err != nil {
		panic(err)
	}

	// apiGroup := r.Group("/api")

}
