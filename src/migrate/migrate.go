package main

import (
	"github.com/FelipeGadelha/go-crud/src/initializers"
	"github.com/FelipeGadelha/go-crud/src/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
