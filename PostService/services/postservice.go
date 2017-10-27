package services


type PostService struct {
	posts  []Post
}


func (p *PostService) GetPosts() []Post {
	pc := PostService{posts: []Post{
		{Id: 1, Title: "A good post", Body: "A very very very long post"},
		{Id: 2, Title: "A bad post", Body: "A very very very short post"},
	}}
	return pc.posts
}