package main

import (
	"fmt"
	"postservice/api"
	"log"
	"net/http"
)

func main() {

	var appName= "post-service"
	fmt.Printf("Staring %v\n", appName)

	StartWebServer("5003")



}
func StartWebServer(port string) {
	router := api.NewRouter()
	http.Handle("/", router)

	log.Println("Starting HTTP server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting the server at port " + port)
		log.Println("Error: " + err.Error())
	}
}