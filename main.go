package main

import (
	"github.com/FelipeGadelha/go-crud/src/controllers"
	"github.com/FelipeGadelha/go-crud/src/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.Create)
	r.GET("/posts", controllers.FindAll)
	r.GET("/posts/:id", controllers.FindById)
	r.PUT("/posts/:id", controllers.Update)
	r.DELETE("/posts/:id", controllers.Remove)
	r.Run()
}
