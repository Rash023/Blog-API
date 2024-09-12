package main

import (
	"go_crud/controllers"
	"go_crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}
func main() {
	r := gin.Default()
	r.POST("/create", controllers.CreatePost)
	r.GET("/view", controllers.ViewPosts)
	r.GET("/read/:id", controllers.ReadPost)
	r.PUT("/update/:id", controllers.UpdatePost)
	r.DELETE("/delete/:id", controllers.DeletePost)
	r.Run()
}
