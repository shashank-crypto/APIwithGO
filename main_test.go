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

