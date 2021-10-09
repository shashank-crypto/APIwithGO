package controller

import (
	"insta/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(c *gin.Context) {
	var user models.Users
	userId := c.Param("userId")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	err := client.Database("instagram").Collection("users").FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		log.Printf("Coun't get the Post")
	}
	c.JSON(http.StatusOK, gin.H{"userId": user.Id, "username": user.Name})
}