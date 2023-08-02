package main

import (
	"github.com/Izzy4999/gin_gorm/initializers"
	"github.com/Izzy4999/gin_gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
