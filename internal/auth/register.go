package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JoaoRafa19/goplaningbackend/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RegisterClient(request UserRegisterBody) error {

	hashedPassword, err := hashPassword(request.Password)

	if err != nil {
		return err
	}

	databaseUser := database.InsertUserModel{Username: request.Username, HashedPassword: hashedPassword}
	return saveClient(databaseUser)
}

func saveClient(user database.InsertUserModel) error {

	localCtx, cancel := context.WithTimeout(context.TODO(), 2000*time.Hour)
	defer cancel()

	respch := make(chan error)

	go func() {

		savedUser, err := database.GetUser(user.Username)
		if savedUser != nil && err == nil {
			respch <- fmt.Errorf("user alredy created")
			return
		}
		if err != nil {
			if error.Error(err) != "user_not_found" {
				panic(err)
			} else {
				log.Println(err)
			}
		}
		respch <- database.Insert(database.UsersCollection, user)
	}()

	for {
		select {
		case resp := <-respch:
			return resp
		case <-localCtx.Done():
			return fmt.Errorf("inserting data tooks too long")
		}
	}

}
