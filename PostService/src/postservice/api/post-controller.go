package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"postservice/data"
	"postservice/services"
)

type PostController struct {
	postService *services.PostService
}

func (pc *PostController) PostController(ps *services.PostService) {
	ps.Init()
	pc.postService = ps
}

func (pc PostController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, pc.postService.GetAll())

}

func (pc *PostController) Add(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	post := data.Post{Body: c.PostForm("body"), Title: c.PostForm("title"), Id: id}
	pc.postService.Add(post)

	// todo: Message doesn't print id well
	c.JSON(200, gin.H{
		"message": "Post with id " + string(id) + " has been added",
	})
}

func (pc *PostController) Edit(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	post := data.Post{Body: c.PostForm("body"), Title: c.PostForm("title"), Id: id}
	pc.postService.Edit(post)
	// todo: Message doesn't print id well
	c.JSON(200, gin.H{
		"message": "Post with id " + string(id) + " has been changed",
	})
}

func (pc *PostController) Remove(c *gin.Context) {
	// todo
}

func (pc *PostController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, pc.postService.Get(id))

}
