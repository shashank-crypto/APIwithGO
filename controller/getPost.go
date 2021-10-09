package controller

import (
	"insta/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPost(c *gin.Context) {
	var post models.Posts
	postId := c.Param("postId")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	err := client.Database("instagram").Collection("posts").FindOne(ctx, bson.M{"_id": postId}).Decode(&post)
	if err != nil {
		log.Printf("Coun't get the Post")
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}