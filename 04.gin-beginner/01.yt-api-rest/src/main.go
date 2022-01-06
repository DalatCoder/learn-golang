package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func HomePage(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "hello world",
	})
}

func PostHomePage(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "post home page",
	})
}

func QueryString(context *gin.Context) {
	name := context.Query("name")
	age := context.Query("age")

	context.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PathParameters(context *gin.Context) {
	name := context.Param("name")
	age := context.Param("age")

	context.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PostWithBody(context *gin.Context) {
	body := context.Request.Body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println(err.Error())
	}

	context.JSON(200, gin.H{
		"message": string(value),
	})
}

func main() {
	r := gin.Default()

	r.GET("/", HomePage)
	r.POST("/", PostHomePage)

	r.GET("/query", QueryString) // query?name=Hieu&age=20
	r.GET("/path/:name/:age", PathParameters) // query/hieu/20

	r.POST("/body", PostWithBody) // query/hieu/20

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
