package handlers

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Serve(ch chan amqp.Delivery) {

	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router, ch)

	// Run the server
	router.Run(":3000") // listen and serve at localhost:8080
}
