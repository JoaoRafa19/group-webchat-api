package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"nhooyr.io/websocket"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	basePath := "/"
	router.LoadHTMLGlob("templates/*")
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Opening routes
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello"})
		})

		v1.GET("/ws", func(c *gin.Context) {
			// c.Header("Content-Type", "text/event-stream")
			// c.Header("Cache-Control", "no-cache")
			// c.Header("Connection", "keep-alive")

			conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
				InsecureSkipVerify: true,
			})

			if err != nil {
				log.Fatal(err)
			}

			for {
				_, msg, err := conn.Read(c)
				if err != nil {
					log.Println(err)
					break
				}
				if string(msg) == "ola" {
					conn.Write(c, websocket.MessageText, []byte("hi"))
				} else {
					conn.Write(c, websocket.MessageText, []byte(string(msg)))
				}
				log.Println(string(msg))
			}
		})
	}
	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

}
