package controllers

import (
	"go_crud/initializers"
	"go_crud/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// get data from request body and then create the post
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)
	post := models.Post{Title: body.Title, Content: body.Content}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": post,
	})
}

func ViewPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
