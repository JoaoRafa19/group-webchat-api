package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JoaoRafa19/goplaningbackend/internal/auth"
	"github.com/gin-gonic/gin"
)

func LoginHandler(context *gin.Context) {

	var req auth.UserLoginRequest

	if err := json.NewDecoder(context.Request.Body).Decode(&req); err != nil {
		context.Error(err)
		return
	}

	otp, err := auth.LoginHandler(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	otpErr := manager.Otps.NewOTP(otp)
	if otpErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	context.JSON(http.StatusAccepted, gin.H{"otp": otp})
}
