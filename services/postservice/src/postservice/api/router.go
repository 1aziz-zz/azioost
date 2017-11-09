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
	postService := services.PostService{}
	postController := PostController{}

	postService.Init(getDbSettings("conf.json"), "posts")

	postController.PostController(&postService)

	for _, route := range postController.GetRoutes() {
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
