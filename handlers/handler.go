package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func handler (c *gin.Context){
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Fatal(err)
	}
	
	 _, msg, err := conn.Read(c) 
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(msg))


}

func Serve() {

	// Initialize router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":3000") // listen and serve at localhost:8080
}
