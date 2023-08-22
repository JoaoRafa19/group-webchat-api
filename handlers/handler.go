package handlers

import (

	"github.com/gin-gonic/gin"
)

func Serve() {

	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080") // listen and serve at localhost:8080
}
