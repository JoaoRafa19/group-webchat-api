package auth

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JoaoRafa19/goplaningbackend/pkg/database"
	"github.com/google/uuid"
)

/*
*

1- A regular HTTP request to Authenticate returns an OneTimePassword (OTP) which can be used to connect to a WebSocket connection.

2- Connect WebSocket, but don’t accept any messages until a special Authentication message with credentials has been sent.
*/
func LoginHandler(req UserLoginRequest) (string, error) {

	//The user authenticates using regular HTTP, an OTP ticket is returned to the user.
	dbUser, err := database.GetUser(req.Username)

	if err != nil && dbUser == nil {
		log.Println(err)
		return "", err
	}
	if err != nil {
		log.Println(err)
		return "", err
	}

	var user database.InsertUserModel
	unmashErr := json.Unmarshal(dbUser, &user)
	if unmashErr != nil {
		return "", unmashErr
	}

	if req.Username == user.Username && checkPasswordHash(req.Password, user.HashedPassword) {
		otp := uuid.NewString()
		return otp, nil
	}

	return "", fmt.Errorf("unauthorised")
}

