package services

import (
	"postservice/data"
)

type PostService struct {
	posts []data.Post
}

func (pc *PostService) Init() {
	pc.posts = []data.Post{
		{Id: 1, Body: "Top", Title: "Title"},
		{Id: 2, Body: "Pot", Title: "Little"},
	}
}

func (pc *PostService) GetAll() []data.Post {
	return pc.posts

}

func (pc *PostService) Add(post data.Post) {
	pc.posts = append(pc.posts, post)
}

func (pc *PostService) Edit(post data.Post) {
	pc.posts = append(pc.posts, post)
}

func (pc *PostService) Remove(comment data.Comment) {
	// todo
}

func (pc *PostService) Get(id int) data.Post {
	for _, element := range pc.posts {
		if element.Id == id {
			return element
		}
	}
	return data.Post{}
}
