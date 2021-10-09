package controller

import (
	"insta/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListPost(c *gin.Context) {
	var postList []models.Posts
	var post models.Posts
	userId := c.Param("userId")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	findOptions := options.Find()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	var perPage int64 = 2

	findOptions.SetSkip((int64(page) - 1) * perPage)
	findOptions.SetLimit(perPage)

	cur, err := client.Database("instagram").Collection("posts").Find(ctx, bson.M{"author": userId}, findOptions)
	if err != nil {
		log.Printf("having some trouble finding the documents")
		defer cur.Close(ctx)
	} else {
		for cur.Next(ctx) {
			err = cur.Decode(&post)
			if err != nil {
				log.Printf("having trouble decoding the cur", err)
				return
			}
			postList = append(postList, post)
		}
	}
	c.JSON(http.StatusOK, gin.H{"postList": postList})
}