package store

import (
	"fmt"
	"github.com/google/uuid"
	"goTinyUrl/shortener"
	"testing"
)

func TestCassandra(t *testing.T) {
	Init()
}

func TestSave(t *testing.T) {
	shortUrl := shortener.GenerateShortLink("http://www.google.com", uuid.New().String())
	err, ok := Save("http://www.google.com", shortUrl, "127.0.0.1")
	fmt.Println(shortUrl, err, ok)
}

func TestGet(t *testing.T) {
	shortUrl := "avRbHRzu"
	originUrl := Get(shortUrl)
	fmt.Println(originUrl)
}
