package main

import (
	"github.com/sgl-soh/golang-prj/initializers"
	"github.com/sgl-soh/golang-prj/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})

}
