package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func TestRoutes(r *gin.RouterGroup, testController *controllers.TestController, authService *services.AuthService) {
	tests := r.Group("/tests")
	{
		// Authenticated routes
		tests.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		tests.Use(middlewares.AdminMiddleware)
		{
			tests.POST("/create", controllers.CreateTestHandler)
			tests.PUT("/:testID", controllers.UpdateTestHandler)
			tests.DELETE("/:testID", controllers.DeleteTestHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		tests.GET("/:testID", controllers.GetTestHandler)
	}
}
