package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Users struct {
	Id string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Posts struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Caption string `json:"caption" bson:"caption"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	Author string `json:"author" bson:"author"`
	Time time.Time `json:"time" bson:"time"`
}


// func connectMongo() *mongo.Client {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err.Error)
// 	}
// 	return client
// }

// func getMongo() {

// 	clientOptions := options.Client().
// 		ApplyURI("mongodb+srv://go_api:get_api@cluster0.opczi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://go_api:get_api@cluster0.opczi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Printf("Failed to create client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func createUser(user *Users) (string, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	result, err := client.Database("instagram").Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Printf("Couldn't create the user")
	}
	uid := result.InsertedID.(string)
	return uid, err
}

func getUser(c *gin.Context) {
	var user Users
	userId := c.Param("userId")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	err := client.Database("instagram").Collection("users").FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		log.Printf("Coun't get the Post")
	}
	c.JSON(http.StatusOK, gin.H{"userId" : user.Id,"username" : user.Name})
}

func handleCreateUser(c *gin.Context) {
	var newUser Users
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := createUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg" : err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id" : id})
}

func handleCreatePost(c *gin.Context) {
	var newPost Posts
	if err := c.ShouldBindJSON(&newPost); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := createPost(&newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg" : err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id" : id})
}

func createPost(post *Posts) (primitive.ObjectID, error) {
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

func getPost(c *gin.Context) {
	var post Posts
	postId := c.Param("postId")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	err := client.Database("instagram").Collection("posts").FindOne(ctx, bson.M{"_id": postId}).Decode(&post)
	if err != nil {
		log.Printf("Coun't get the Post")
	}
	c.JSON(http.StatusOK, gin.H{"post" : post})
}

func listPost(c *gin.Context) {
	var postList []Posts
	var post Posts
	userId := c.Param("userId")
	client, ctx, cancel :=getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	cur, err := client.Database("instagram").Collection("posts").Find(ctx, bson.M{"author" : userId})
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
	c.JSON(http.StatusOK, gin.H{"postList" : postList})
}

func main() {
	router := gin.Default()
	router.GET("/users/:userId" , getUser)
	router.POST("/users" , handleCreateUser)
	router.POST("/posts" , handleCreatePost)
	router.GET("/posts/:postId" , getPost)
	router.GET("/posts/users/:userId" , listPost)
	router.Run()
}
