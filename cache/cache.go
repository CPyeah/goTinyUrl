package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

var Ctx context.Context
var Rdb *redis.Client

const cachePrefix = "tinyURL:"

func init() {
	Ctx = context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       1,                                  // use default DB
	})

	pong, err := Rdb.Ping(Ctx).Result()
	if err == nil {
		fmt.Println("connect success", pong)
	} else {
		fmt.Println("connect fail")
		panic(err)
	}
}

func Sava(shortUrl string, originUrl string) {
	Rdb.Set(Ctx, cachePrefix+shortUrl, originUrl, time.Hour*6)
}

func Get(shortUrl string) string {
	result, err := Rdb.Get(Ctx, cachePrefix+shortUrl).Result()
	if err != nil {
		log.Default().Println(err)
		return ""
	}
	fmt.Println("Cache hit")
	return result
}
