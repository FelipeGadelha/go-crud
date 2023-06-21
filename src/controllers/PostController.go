package controllers

import (
	"net/http"

	"github.com/FelipeGadelha/go-crud/src/initializers"
	"github.com/FelipeGadelha/go-crud/src/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// log.Println("Create Post Method")
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func FindAll(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	// c.JSON(http.StatusOK, posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func FindById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.Find(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.Find(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Remove(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id) // soft delete

	c.Status(204)
}
