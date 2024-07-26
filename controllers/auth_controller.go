package controllers

import (
	"net/http"
	"time"

	"github.com/0xSumeet/go_app/database"
	"github.com/0xSumeet/go_app/models"
	"github.com/0xSumeet/go_app/utils"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: req.Username, Password: req.Password, Role: req.Role}
	if err := user.RegisterUser(database.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: req.Username, Password: req.Password}
	if err := user.AuthenticateUser(database.DB); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

// LogoutHandler deletes the JWT cookie
func LogoutHandler(c *gin.Context) {
	// Create a cookie with the same name as the JWT token cookie
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "",
		Expires: time.Now().
			Add(-time.Hour),
		// Set expiration time to the past to delete the cookie
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
