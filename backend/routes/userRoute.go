package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController *controllers.UserController, authService *services.AuthService) {
	// Authenticated routes for all users
	users := r.Group("/users")
	{
		// Use the AuthMiddleware for all routes within this group
		users.Use(middlewares.AuthMiddleware(authService))

		// Routes accessible by authenticated users
		users.PUT("/:id", userController.EditUser)

		// Admin-only routes
		users.Use(middlewares.AdminMiddleware)
		{
			// users.POST("/", userController.CreateUser)
			users.DELETE("/:id", userController.DeleteUser)
			users.PUT("/:id/role", userController.ChangeUserRole)
		}
	}
}
