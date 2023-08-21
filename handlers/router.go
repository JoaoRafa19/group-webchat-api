package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func initializeRoutes(router *gin.Engine, ch chan amqp.Delivery) {
	//Initialize handler
	basePath := "/"
	router.LoadHTMLGlob("templates/*")
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Opening routes
		v1.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.tpml", gin.H{"title": "SSE", "url": "/sse"})
		})

		v1.GET("/sse", func (c * gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			for m := range ch {
				fmt.Fprintf(c.Writer, "event: message\n")
				fmt.Fprintf(c.Writer, "data: %s\n\n", m.Body)
				c.Writer.Flush()
			}
		})

	}
	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

}
