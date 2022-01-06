package main

import (
	"github.com/dalatcoder/yt-beginner/httpd/handler"
	"github.com/dalatcoder/yt-beginner/platform/newfeeds"
	"github.com/gin-gonic/gin"
)

func main()  {
	feed := newfeeds.New()
	// db := db.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
