package main

import (
	"go_user_auth/internal/config"
	"go_user_auth/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Init Database
	db := config.NewDatabase()

	// 2. Run Migrations (Create Tables automatically)
	db.AutoMigrate()

	// 3. Init Layers (Dependency Injection)
	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	// 4. Setup Router
	r := gin.Default()
	r.Static("/uploads", "./uploads") // Serve images publicly

	// 5. Routes
	r.POST("/api/register", handler.Register)
	r.POST("/api/login", handler.Login)

	// Protected Routes (Require Token)
	// protected := r.Group("/api/user")
	// protected.Use(user.AuthMiddleware())
	// {
	// 	protected.POST("/change-password", handler.ChangePassword)
	// 	protected.POST("/profile-image", handler.UploadImage)
	// 	protected.POST("/logout", handler.Logout)
	// }

	// 6. Start Server
	r.Run(":8080")
}
