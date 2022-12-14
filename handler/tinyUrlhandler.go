package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"goTinyUrl/analyst"
	"goTinyUrl/shortener"
	"goTinyUrl/store"
	"net/http"
)

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, uuid.New().String())
	err, ok := store.Save(creationRequest.LongUrl, shortUrl, c.ClientIP())

	if !ok {
		shortUrl = shortener.GenerateShortLink(creationRequest.LongUrl, uuid.New().String())
		err, ok = store.Save(creationRequest.LongUrl, shortUrl, c.ClientIP())
		if !ok {
			shortUrl = shortener.GenerateShortLink(creationRequest.LongUrl, uuid.New().String())
			err, ok = store.Save(creationRequest.LongUrl, shortUrl, c.ClientIP())
		}
	}

	if err != nil || !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "short url created error",
		})
		return
	}

	host := "http://localhost:8080/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	originUrl, err := store.Get(shortUrl)
	if len(originUrl) == 0 || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "short url error",
		})
		return
	}
	go analyst.Analysis(shortUrl, originUrl, c.ClientIP())
	c.Redirect(302, originUrl)
}
