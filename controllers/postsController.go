package controllers

import (
	"github.com/Izzy4999/gin_gorm/initializers"
	"github.com/Izzy4999/gin_gorm/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// create post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return the post
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get the post
	var post []models.Post
	initializers.DB.Find(&post, id)

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// get data of req.body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)

	// find the post
	var post []models.Post
	initializers.DB.Find(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get the id of  post
	id := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// repsonse
	c.Status(200)
}
