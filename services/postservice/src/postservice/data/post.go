package data

type Post struct {
	Id    int    `api:"id"`
	Title string `api:"title"`
	Body string `api:"body"`
}
