package controllers

import (
	"go-curd/helper"
	"go-curd/initializer"
	"go-curd/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	// get req body
	data := helper.Body
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create record
	post := models.Post{Title: data.Title, Body: data.Body}
	result := initializer.DB.Create(&post)

	// error case cover
	if result.Error != nil {
		c.JSON(400, gin.H{"message": "could not create record"})
		return
	}

	// return value
	c.JSON(200, gin.H{"post": post})
}

func GetPosts(c *gin.Context) {
	// Get posts
	var posts []models.Post
	initializer.DB.Find(&posts)

	// return posts
	c.JSON(200, gin.H{"posts": posts})

}

func GetPostById(c *gin.Context) {
	// Get id of post
	id := c.Param("id")
	var post models.Post

	// find data with id
	data := initializer.DB.First(&post, id)

	// cover error case
	if data.Error != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	// return data
	c.JSON(200, gin.H{"post": post})
}

func UpdatePostById(c *gin.Context) {
	// Get id of post
	id := c.Param("id")
	var post models.Post

	// Declare a new variable for the request body
	var data struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	// Bind body of the request
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find post by id in the database
	findPost := initializer.DB.First(&post, id)

	// Cover error case
	if findPost.Error != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	// Update post
	initializer.DB.Model(&post).Updates(models.Post{Title: data.Title, Body: data.Body})

	// Return the updated post
	c.JSON(200, gin.H{"post": post})
}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")

	findPost := initializer.DB.Delete(&models.Post{}, id)

	if findPost.Error != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(200, gin.H{"message": "Post has been Deleted!"})

}
