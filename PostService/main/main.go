package main

import (
	"github.com/gin-gonic/gin"
	"goPlay/api"
)

func main() {
	router := gin.Default()

	router.GET("/", api.GetPosts)
	router.PUT("/edit", api.EditPost)
	/*	router.POST("/remove", removePost)
		router.POST("/create", addPost)*/

	router.Run() // listen and serve on 0.0.0.0:8080
}
