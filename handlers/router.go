package handlers

import (
	"github.com/JoaoRafa19/goplaningbackend/session"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

var (
	rooms map[string]map[*session.Manager]bool = make(map[string]map[*session.Manager]bool)
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	basePath := "/"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Opening routes
		v1.GET("/", CreateRooms)

		v1.GET("/clients", GetClients)

		v1.GET("/rooms", GetRooms)

		v1.GET("/ws/:room_id", ConnectRoom)
	}
	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

}
