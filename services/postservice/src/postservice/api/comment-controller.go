package api

import (
	"postservice/services"
	"postservice/data"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"net/http"
	"encoding/json"
)

type CommentController struct {
	commentService *services.CommentService
}

func (cc *CommentController) CommentController(cs *services.CommentService) {
	cc.commentService = cs
}

func (cc CommentController) GetAll(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(cc.getJson(cc.commentService.GetAll()))
}

func (cc *CommentController) Add(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	body := request.Form.Get("body")
	title := request.Form.Get("title")

	result := "ERROR: please provide comment properties."

	if body != "" && title != "" {
		if cc.commentService.CheckCommentBody(body) {
			comment := data.Comment{Body: request.Form.Get("body"), Title: request.Form.Get("title")}
			cc.commentService.Add(&comment)
			result = "OK"
		}

	}
	w.Write([]byte("{\"result\":\"" + result + "\"}"))

}

func (cc *CommentController) Edit(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := request.Form.Get("id")
	title := request.Form.Get("title")
	body := request.Form.Get("body")
	var result string
	if id != "" && title != "" || body != "" {
		if bson.IsObjectIdHex(id) {
			comment := data.Comment{Body: body, Title: title, Id: bson.ObjectIdHex(id)}
			cc.commentService.Edit(&comment)
			result = "OK"
		} else {
			result = "ERROR: Id is not correct."
		}
	} else {
		result = "ERROR: please provide comment properties."

	}
	w.Write([]byte("{\"result\":\"" + result + "\"}"))

}

func (cc *CommentController) Remove(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := request.Form.Get("id")
	result := "OK"

	if id != "" {
		if bson.IsObjectIdHex(id) {
			if cc.commentService.CommentExists(id) {
				cc.commentService.Remove(cc.commentService.Get(bson.ObjectIdHex(id)))
			} else {
				result = "ERROR: No comment found to remove."
			}
		} else {
			result = "ERROR: Id is not correct."
		}
	} else {
		result = "ERROR: Id is undefined."
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{\"result\":\"" + result + "\"}"))
}

func (cc *CommentController) Get(w http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	id := strings.TrimPrefix(request.URL.Path, "/comment/")
	result := "OK"
	var foundComment []byte
	if bson.IsObjectIdHex(id) {
		oid := bson.ObjectIdHex(id)
		if comment := cc.commentService.Get(oid); (data.Comment{}) != comment {
			foundComment = cc.getJson(comment)
		} else {
			result = "ERROR: Id is not found."
		}
	} else {
		result = "ERROR: Id is not correct."
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if foundComment == nil {
		w.Write([]byte("{\"result\":\"" + result + "\"}"))
	} else {
		w.Write(foundComment)
	}

}
func (cc *CommentController) getJson(data interface{}) []byte {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return jData
}
