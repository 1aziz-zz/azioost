package data

type Comment struct {
	Id   int    `api:"id"`
	Body string `api:"body"`
	Post Post   `api:"Post"`
}
