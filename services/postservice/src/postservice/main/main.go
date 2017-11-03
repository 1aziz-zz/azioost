package main

import (
	"github.com/gin-gonic/gin"
	"postservice/api"
	"postservice/services"
	"fmt"
	"os"
	"encoding/json"
)

func main() {

	router := gin.Default()
	postService := services.PostService{}
	postController := api.PostController{}

	postService.Init(getDbSettings("conf.json"), "posts")
	postController.PostController(&postService)

	router.GET("/", postController.GetAll)
	router.PUT("/edit", postController.Edit)
	router.POST("/remove", postController.Remove)
	router.POST("/create", postController.Add)
	router.GET("/id/:id/", postController.Get)

	router.Run()
}
func getDbSettings(file string) services.DbService {
	var config services.DbService
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
