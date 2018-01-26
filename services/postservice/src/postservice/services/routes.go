package services

import (
	"net/http"
	"postservice/api"
)

type Route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
	pc                    api.PostController
	cc                    api.CommentController
}

func init() {

}

func (route *Route) GetRoutes() []Route {
	type Routes []Route

	return Routes{

		Route{
			Name:        "GetAll",
			Method:      "GET",
			Pattern:     "/posts/",
			HandlerFunc: route.pc.GetAll,
		}, Route{
			Name:        "Get",
			Method:      "GET",
			Pattern:     "/post/{postId}",
			HandlerFunc: route.pc.Get,
		}, Route{
			Name:        "Add",
			Method:      "POST",
			Pattern:     "/post/add",
			HandlerFunc: route.pc.Add,
		}, Route{
			Name:        "Edit",
			Method:      "POST",
			Pattern:     "/post/edit",
			HandlerFunc: route.pc.Edit,
		}, Route{
			Name:        "Remove",
			Method:      "POST",
			Pattern:     "/post/remove",
			HandlerFunc: route.pc.Remove,
		}, Route{Name: "GetAll",
			Method: "GET",
			Pattern: "/comments/",
			HandlerFunc: route.cc.GetAll,
		}, Route{
			Name:        "Get",
			Method:      "GET",
			Pattern:     "/comment/{commentId}",
			HandlerFunc: route.cc.Get,
		}, Route{
			Name:        "Add",
			Method:      "POST",
			Pattern:     "/comment/add",
			HandlerFunc: route.cc.Add,
		},
		Route{
			Name:        "Edit",
			Method:      "POST",
			Pattern:     "/comment/edit",
			HandlerFunc: route.cc.Edit,
		},
		Route{
			Name:        "Remove",
			Method:      "POST",
			Pattern:     "/comment/remove",
			HandlerFunc: route.cc.Remove,
		},
	}
}
