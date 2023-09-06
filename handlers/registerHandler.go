package handlers

import (
	"github.com/JoaoRafa19/goplaningbackend/client/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	auth.RegisterClient(c)
}