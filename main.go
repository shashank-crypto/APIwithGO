package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Users struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Posts struct {
	Id primitive.ObjectID `json:"id"`
	Caption string `json:"caption"`
	ImageUrl string `json:"image_url"`
	Time primitive.Timestamp `json:"time"`
}



func getUser(c *gin.Context) {
	id := c.Param("userId")
	c.IndentedJSON(http.StatusOK , id)
}

func createUser(c *gin.Context) {

}

func createPost(c *gin.Context) {

}

func getPost(c *gin.Context) {

}

func listPost(c *gin.Context) {

}

func main() {
	router := gin.Default()
	router.GET("/users/:userId" , getUser)
	router.POST("/users" , createUser)
	router.POST("/posts" , createPost)
	router.GET("/posts/:postId" , getPost)
	router.GET("/posts/users/:userId" , listPost)
	router.Run()
}
