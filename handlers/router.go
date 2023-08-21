package handlers

import (
	"fmt"
	"net/http"
	"sse/pkg/rabbitmq"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		v1.GET("/sse", func(c *gin.Context) {
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

	v2 := router.Group("/v2")
	{
		v2.GET("/newroom", func(c *gin.Context) {
			room := uuid.New()

			c.JSON(http.StatusCreated, gin.H{"message": "room created", "id": room})
		})
		v2.GET("/:room", func(c *gin.Context) {
			room := c.Params.ByName("room")
			if room == "" {

				c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid room id"})
			}
			conn, err := rabbitmq.OpenChannel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			}
			conn.Close()
			c.JSON(200, gin.H{})
		})
	}

	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

}
