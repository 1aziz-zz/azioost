package api

import (
	"net/http"
	"postservice/data"
	"postservice/services"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type PostController struct {
	postService *services.PostService
}

func (pc *PostController) GetRoutes() []services.Route {
	type Routes []services.Route

	return Routes{

		services.Route{
			Name:        "GetAll",
			Method:      "GET",
			Pattern:     "/posts/",
			HandlerFunc: pc.GetAll,
		}, services.Route{
			Name:        "Get",
			Method:      "GET",
			Pattern:     "/posts/{postId}",
			HandlerFunc: pc.Get,
		}, services.Route{
			Name:        "Add",
			Method:      "POST",
			Pattern:     "/posts/add",
			HandlerFunc: pc.Add,
		},
		services.Route{
			Name:        "Remove",
			Method:      "POST",
			Pattern:     "/posts/remove",
			HandlerFunc: pc.Remove,
		},
	}

}
func (pc *PostController) getJson(data interface{}) []byte {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jData
}

func (pc *PostController) PostController(ps *services.PostService) {
	pc.postService = ps
}

func (pc PostController) GetAll(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(pc.getJson(pc.postService.GetAll()))
}

func (pc *PostController) Add(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	post := data.Post{Body: request.Form.Get("body"), Title: request.Form.Get("title")}

	pc.postService.Add(&post)
	w.Write([]byte("{\"result\":\"OK\"}"))

}

func (pc *PostController) Edit(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	id := request.Form.Get("id")
	post := data.Post{Body: request.Form.Get("body"), Title: request.Form.Get("title"), Id: bson.ObjectIdHex(id)}
	pc.postService.Edit(&post)

	w.Write([]byte("{\"result\":\"OK\"}"))

}

func (pc *PostController) Remove(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := request.Form.Get("id")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pc.postService.Remove(pc.postService.Get(bson.ObjectIdHex(id)))
	w.Write([]byte("{\"result\":\"OK\"}"))
}

func (pc *PostController) Get(w http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	id := strings.TrimPrefix(request.URL.Path, "/posts/")

	if bson.IsObjectIdHex(id) {
		oid := bson.ObjectIdHex(id)
		if post := pc.postService.Get(oid); (data.Post{}) != post {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(pc.getJson(post))
		}

	} else {
		w.WriteHeader(404)
	}
}
