package api

import (
	"postservice/services"
	"postservice/data"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"net/http"
	"encoding/json"
)

type PostController struct {
	postService *services.PostService
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
	body := request.Form.Get("body")
	title := request.Form.Get("title")

	result := "ERROR: please provide post properties."

	if body != "" && title != "" {
		if pc.postService.CheckPostBody(body) {
			post := data.Post{Body: request.Form.Get("body"), Title: request.Form.Get("title")}
			pc.postService.Add(&post)
			result = "OK"
		}

	}
	w.Write([]byte("{\"result\":\"" + result + "\"}"))

}

func (pc *PostController) Edit(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := request.Form.Get("id")
	title := request.Form.Get("title")
	body := request.Form.Get("body")
	var result string
	if id != "" && title != "" || body != "" {
		if bson.IsObjectIdHex(id) {
			post := data.Post{Body: body, Title: title, Id: bson.ObjectIdHex(id)}
			pc.postService.Edit(&post)
			result = "OK"
		} else {
			result = "ERROR: Id was not correct."
		}
	} else {
		result = "ERROR: please provide post properties."

	}
	w.Write([]byte("{\"result\":\"" + result + "\"}"))

}

func (pc *PostController) Remove(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := request.Form.Get("id")
	result := "OK"

	if id != "" {
		if bson.IsObjectIdHex(id) {
			if pc.postService.PostExists(id) {
				pc.postService.Remove(pc.postService.Get(bson.ObjectIdHex(id)))
			} else {
				result = "ERROR: No posts found to remove."
			}
		} else {
			result = "ERROR: Id was not correct."
		}
	} else {
		result = "ERROR: Id was undefined."
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{\"result\":\"" + result + "\"}"))
}

func (pc *PostController) Get(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := strings.TrimPrefix(request.URL.Path, "/post/")
	result := "OK"
	var foundPost []byte
	if bson.IsObjectIdHex(id) {
		oid := bson.ObjectIdHex(id)
		if post := pc.postService.Get(oid); post.Id != "" {
			foundPost = pc.getJson(post)
		} else {
			result = "ERROR: Id was not found."
		}
	} else {
		result = "ERROR: Id was not correct."
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if foundPost == nil {
		w.Write([]byte("{\"result\":\"" + result + "\"}"))
	} else {
		w.Write(foundPost)
	}

}
func (pc *PostController) getJson(data interface{}) []byte {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jData
}
