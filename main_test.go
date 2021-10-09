package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGetUserEndpoint(t *testing.T) {
	router := gin.Default()
	request, _ := http.NewRequest("GET","/users/sha2shank2", nil) 
	response := httptest.NewRecorder()
	router.ServeHTTP(response , request)
	assert.Equal(t, 200, response.Code, "Ok response exprected")
}

func TestGetPostEndpoint(t * testing.T){
	router := gin.Default()
	request, _ := http.NewRequest("GET","/posts/6161d63d42887d58b4e0c748", nil) 
	response := httptest.NewRecorder()
	router.ServeHTTP(response , request)
	assert.Equal(t, 200, response.Code, "Ok response exprected")
}

func TestpostPostEndpoint(t * testing.T){
	router := gin.Default()
	request, _ := http.NewRequest("POST","/posts", {
		"caption" : "sixth Post",
		"imageUrl" : "https://miro.medium.com/max/1132/0*Fbely1We1Tmhfnag",
		"author" : "sha2shank2"
	}) 
	response := httptest.NewRecorder()
	router.ServeHTTP(response , request)
	assert.Equal(t, 200, response.Code, "Ok response exprected")
}

func TestPostUserEndpoint(t * testing.T){
	router := gin.Default()
	request, _ := http.NewRequest("POST","/users", {
		"id" : "newUser",
		"name" : "shashank k",
		"email" : "shas@gmail.com",
		"password" : "appointy"
	}) 
	response := httptest.NewRecorder()
	router.ServeHTTP(response , request)
	assert.Equal(t, 200, response.Code, "Ok response exprected")
}

func TestGetPostListEndpoint(t * testing.T){
	router := gin.Default()
	request, _ := http.NewRequest("GET","/posts/users/sha2shank2", nil) 
	response := httptest.NewRecorder()
	router.ServeHTTP(response , request)
	assert.Equal(t, 200, response.Code, "Ok response exprected")
}
