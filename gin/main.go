package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.GET("/upload", func(c *gin.Context) {
		c.String(http.StatusOK, "please sign up")
	})
	router.Static("/ex", "./ex")
	router.Run("0.0.0.0:3000")
	//router.NoRoute("./page404")
}
