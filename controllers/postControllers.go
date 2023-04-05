package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sgl-soh/golang-prj/initializers"
	"github.com/sgl-soh/golang-prj/models"
)

func PostsList(c *gin.Context) {
	// Get all records
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the posts list.",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Find the post with that id
	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The server has not found anything matching the Request-URI",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostCreator(c *gin.Context) {
	// Get the data from req
	var data struct {
		Title string
		Body  string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the data from request.",
		})
		return
	}

	// Create a post
	post := models.Post{Title: data.Title, Body: data.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create the post.",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the data from request
	var data struct {
		Title string
		Body  string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the data from request.",
		})
		return
	}

	// Find related post
	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The server has not found anything matching the Request-URI",
		})
		return
	}

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: data.Title, Body: data.Body})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find related post
	var post models.Post
	initializers.DB.First(&post, id)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The server has not found anything matching the Request-URI",
		})
		return
	}

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	message := fmt.Sprintf("Post with id = %v is deleted successfully", id)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
