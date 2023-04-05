package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sgl-soh/golang-prj/initializers"
	"github.com/sgl-soh/golang-prj/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(c *gin.Context) {
	// Get the data from req
	var data struct {
		Email    string
		Password string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the data from request.",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password.",
		})
	}

	// Create a user
	user := models.User{Email: data.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user.",
		})
	}
	// Response
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UserLogin(c *gin.Context) {
	// Get the info from req
	var data struct {
		Email    string
		Password string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get the data from request.",
		})
	}

	// Look up the user from the email
	var user models.User
	initializers.DB.First(&user, "email = ?", data.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or Password is incorrect.",
		})
	}

	// Compare the emails password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or Password is incorrect.",
		})
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate the tokenString.",
		})
	}

	// Send cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 10, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func UserValidate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
