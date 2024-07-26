package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/0xSumeet/go_app/controllers"
	"github.com/0xSumeet/go_app/database"
	"github.com/0xSumeet/go_app/middlewares"
	"github.com/0xSumeet/go_app/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB(&models.User{}, &models.Permission{})

	r := gin.Default()

	// Register handlers
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)
	r.POST("/logout", controllers.LogoutHandler)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET(
			"/customer-management",
			middlewares.PermissionMiddleware("Sales", "read/write"),
			controllers.CustomerManagement,
		)
		auth.GET(
			"/billing-management",
			middlewares.PermissionMiddleware("Sales", "read/write"),
			controllers.BillingManagement,
		)
		auth.GET(
			"/payroll-management",
			middlewares.PermissionMiddleware("HR", "read/write"),
			controllers.PayrollManagement,
		)
		auth.GET(
			"/user-management",
			middlewares.PermissionMiddleware("Administrator", "read/write"),
			controllers.UserManagement,
		)
	}

	r.Run(":8080")
}
