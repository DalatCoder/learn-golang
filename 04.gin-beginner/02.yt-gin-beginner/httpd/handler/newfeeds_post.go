package handler

import (
	"github.com/dalatcoder/yt-beginner/platform/newfeeds"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newsFeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newfeeds.Adder) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestBody := newsFeedPostRequest{}
		err := context.Bind(&requestBody)
		if err != nil {
			return
		}

		item := newfeeds.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}

		feed.Add(item)

		context.JSON(http.StatusCreated, gin.H{
			"status": "success",
			"data": item,
		})
	}
}
