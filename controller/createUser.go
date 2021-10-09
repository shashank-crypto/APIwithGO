package controller

import (
	"insta/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func createUser(user *models.Users) (string, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	password := []byte(user.Password)
	newPassword, _ := bcrypt.GenerateFromPassword(password, 10)
	user.Password = string(newPassword)
	result, err := client.Database("instagram").Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Printf("Couldn't create the user")
	}
	uid := result.InsertedID.(string)
	return uid, err
}