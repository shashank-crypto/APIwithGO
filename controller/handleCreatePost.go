package controller

import (
	"insta/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreatePost(c *gin.Context) {
	var newPost models.Posts
	if err := c.ShouldBindJSON(&newPost); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := createPost(&newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}