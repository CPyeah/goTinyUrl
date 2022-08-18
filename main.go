package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goTinyUrl/handler"
	"net/http"
)

func main() {
	router := gin.Default()

	// get http://localhost:8080/
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to TinyUrl")
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	fmt.Println("http://localhost:8080")
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
