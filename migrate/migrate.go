package main

import (
	"go_crud/initializers"
	"go_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
