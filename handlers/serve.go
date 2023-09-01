package handlers

import (
	"errors"

	"github.com/JoaoRafa19/goplaningbackend/client"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

var manager * client.Manager

func Serve() {

	// Initialize router
	router := gin.Default()


	basePath := "/"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Opening routes
		v1.GET("/", CreateRooms)

		v1.GET("/clients", GetClients)

		v1.GET("/ws/:room_id", ConnectRoom)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithError(404, c.Error(errors.New("route not found")))
	})
	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

	// Run the server
	router.Run(":8080") // listen and serve at localhost:8080
}
