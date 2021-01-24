package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"log"
	"os"
)

type session struct {
	*redis.Client
	FrontendHost, BackendHost string
}

func main() {

	//localTest()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := client.Ping().Result()
	log.Println("Got pong:", pong)
	if err != nil {
		log.Fatal("Couldn't connect to Redis: ", err.Error())
	}

	ses := session{Client: client, FrontendHost: os.Getenv("FRONTEND_HOST"), BackendHost: os.Getenv("BACKEND_HOST")}

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/setbot", ses.setBot)
	r.POST("/telegram/:token", ses.telegram)

	r.Run("0.0.0.0:" + os.Getenv("PORT"))

}
