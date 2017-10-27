package api

import (
	"github.com/gin-gonic/gin"
	"goPlay/data"
	"net/http"
	"strconv"
)

var posts []data.Post

func init() {
	posts = []data.Post{
		data.Post{Id: 1, Body: "Top", Title: "Title"},
		data.Post{Id: 2, Body: "Pot", Title: "Little"},
	}
}

func GetPosts(c *gin.Context) {
	//todo: error handling
	c.JSON(http.StatusOK, posts)

}

func EditPost(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	post := data.Post{Body: c.PostForm("body"), Title: c.PostForm("title"), Id: id}
	posts = append(posts, post)

	// todo: Message doesn't print id well
	c.JSON(200, gin.H{
		"message": "Post with id " + string(id) + " has been added",
	})
}

func RemovePost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func AddPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
