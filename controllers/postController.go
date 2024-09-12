package controllers

import (
	"go_crud/initializers"
	"go_crud/models"

	"github.com/gin-gonic/gin"
)

//controller to create a new post

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

// controller to view all the posts

func ViewPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// controller to read posts by id
func ReadPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

//controller to update the post

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	//updating the post in the database
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Content: body.Content})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
