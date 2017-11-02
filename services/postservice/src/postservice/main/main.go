package main

import (
	"github.com/gin-gonic/gin"
	"postservice/api"
	"postservice/services"
)

func main() {

	router := gin.Default()
	postService := services.PostService{}
	postController := api.PostController{}

	postController.PostController(&postService)

	router.GET("/", postController.GetAll)
	router.PUT("/edit", postController.Edit)
	router.POST("/remove", postController.Remove)
	router.POST("/create", postController.Add)
	router.GET("/id/:id/", postController.Get)

	router.Run()
}
