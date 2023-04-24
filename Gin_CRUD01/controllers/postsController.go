package controllers

import (
	"CRUD/initializers"
	"CRUD/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // 데이터 생성

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts) //모두 조회

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	//Get id off
	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id) //id 값에 해당 되는 것만 조회

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id) //id 값에 해당 되는 것만 조회

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
