package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoaoRafa19/goplaningbackend/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(context *gin.Context) {

	var body auth.UserRegisterBody

	if err := json.NewDecoder(context.Request.Body).Decode(&body); err != nil {
		ginError := context.AbortWithError(http.StatusBadRequest, err)
		if ginError != nil {
			log.Println(ginError)
			return
		}
		return
	}

	if body.Password == "" || body.Username == "" {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := auth.RegisterClient(auth.UserRegisterBody{Username: body.Username, Password: body.Password})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user_created"})

}
