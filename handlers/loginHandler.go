package handlers

import (
	"github.com/JoaoRafa19/goplaningbackend/client"
	"github.com/gin-gonic/gin"
)

func LoginHandler(context *gin.Context) {
	if manager == nil {
		manager = client.CreateManager(context)
	}
	manager.LoginHandler(context)
}