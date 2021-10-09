package controller

import (
	"insta/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createPost(post *models.Posts) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	post.Time = time.Now()
	post.Id = primitive.NewObjectID()
	result, err := client.Database("instagram").Collection("posts").InsertOne(ctx, post)
	if err != nil {
		log.Printf("Couldn't create the Post", err)
	}
	uid := result.InsertedID.(primitive.ObjectID)
	return uid, err
}