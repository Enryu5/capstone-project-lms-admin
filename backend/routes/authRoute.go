package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {
	// Public routes
	router.POST("/login", authController.LoginHandler)
	router.POST("/register", authController.RegisterHandler)
	router.POST("/logout", authController.LogoutHandler)
}
