package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sgl-soh/golang-prj/controllers"
	"github.com/sgl-soh/golang-prj/initializers"
	"github.com/sgl-soh/golang-prj/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// user
	r.POST("/user/signup", controllers.UserSignUp)
	r.POST("/user/login", controllers.UserLogin)
	r.GET("/user/validate", middleware.RequireAuth, controllers.UserValidate)

	// post
	r.GET("/posts", controllers.PostsList)
	r.GET("/post/:id", controllers.PostShow)
	r.POST("/post", controllers.PostCreator)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
