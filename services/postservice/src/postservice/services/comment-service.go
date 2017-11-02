package services

import "postservice/data"

type CommentController struct {
	comments []data.Comment
}

func (cp *CommentController) Init() {
	cp.comments = []data.Comment{
		{Id: 1, Body: "Top comment", Post: data.Post{}},
		{Id: 2, Body: "Best comment", Post: data.Post{}},
	}
}
func (cp *CommentController) GetAll() []data.Comment {
	return cp.comments

}

func (cp *CommentController) Add(comment data.Comment) {
	cp.comments = append(cp.comments, comment)
}

func (cp *CommentController) Edit(comment data.Comment) {
	cp.comments = append(cp.comments, comment)
}

func (cp *CommentController) Remove(comment data.Comment) {
	// todo
}

func (cp *CommentController) Get(id int) data.Comment {
	for _, comment := range cp.comments {
		if comment.Id == id {
			return comment
		}
	}
	return data.Comment{}
}
