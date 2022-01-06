package handler

import (
	"github.com/dalatcoder/yt-beginner/platform/newfeeds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewsfeedGet(feed newfeeds.Getter) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": feed.GetAll(),
		})
	}
}
