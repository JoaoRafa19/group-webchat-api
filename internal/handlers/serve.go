package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/JoaoRafa19/goplaningbackend/internal/client"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

var manager *client.Manager

func Serve() {

	// Initialize router
	router := gin.Default()

	manager = client.CreateManager(context.Background() )

	basePath := "/"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Opening routes
		v1.GET("/", CreateRooms)

		v1.POST("/register", RegisterHandler)

		v1.POST("/login", LoginHandler)

		v1.GET("/clients", GetClients)

		v1.GET("/ws/:room_id", ConnectRoom)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithError(404, c.Error(errors.New("route not found")))
	})
	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))
	

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Listen error", err)
	}

	// Run the server with TLS
	// err := router.RunTLS(":8080", "server.crt", "server.key") // listen and serve at localhost:8080
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
