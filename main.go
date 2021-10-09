package main

import (
	"insta/controller"
	// "insta/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users/:userId" , controller.GetUser)
	router.POST("/users" , controller.HandleCreateUser)
	router.POST("/posts" , controller.HandleCreatePost)
	router.GET("/posts/:postId" , controller.GetPost)
	router.GET("/posts/users/:userId" , controller.ListPost)
	// router.GET("/get", getBody)
	router.Run()
}
