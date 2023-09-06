package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JoaoRafa19/goplaningbackend/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DatabaseUser struct {
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RegisterClient(context *gin.Context) {

	var body RegisterBody

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

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "bad_password"})
		return
	}

	databaseUser := DatabaseUser{Username: body.Username, HashedPassword: hashedPassword}

	databaseError := saveClient(databaseUser)
	if databaseError != nil {
		log.Print(databaseError)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "database_error", "error": databaseError.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "user_created"})
}

func saveClient(user DatabaseUser) error {

	type  Response struct {
		value *mongo.InsertOneResult
		err error
	}

	localCtx, cancel := context.WithTimeout(context.TODO(), 2000 * time.Hour)
	defer cancel()

	respch := make(chan Response)


	go func () {
		
		// user ,err:= database.GetUser(user.Username) 
		// if user != nil && err != nil {

		// }

		// err := database.Insert("usuarios",user)
		// respch <- Response{
		// 	value: nil,
		// 	err: err,
		// }
	}()
    
	for  {
		select{
		case resp:= <-respch:
			log.Println("User created", resp)
			return resp.err
		case <- localCtx.Done():
			return fmt.Errorf("inserting data tooks too long")	
		}
	}

}



func checkUser(user DatabaseUser) (bool, error) {

	databaseUser, err := database.GetUser(user.Username)

	if err != nil {
		return false, err
	}

	var ret DatabaseUser
	eror := json.Unmarshal(databaseUser, &ret)
	if eror != nil {
		log.Println("Unmarshal error", eror)
		return false, nil
	}

	if user.HashedPassword == ret.HashedPassword {
		return true, nil
	} else {
		return false, nil
	}

}
