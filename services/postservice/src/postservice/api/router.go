package api

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"encoding/json"
	"postservice/services"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	dbSettings := getDbSettings("conf.json")

	postController := PostController{}
	commentController := CommentController{}

	postService := services.PostService{}
	postService.PostService(dbSettings, "posts")

	commentService := services.CommentService{}
	commentService.CommentService(dbSettings, "comments")

	postController.postService = &postService
	commentController.commentService = &commentService

	services.Routes{}

	for _, route := range services..GetRoutes() {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router

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
