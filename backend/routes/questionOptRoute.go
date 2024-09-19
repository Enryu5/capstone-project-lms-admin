package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func QuestionOptRoutes(r *gin.RouterGroup, questionOptController *controllers.QuestionOptController, authService *services.AuthService) {
	questionOpts := r.Group("/questionOpts")
	{
		// Authenticated routes
		questionOpts.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		questionOpts.Use(middlewares.AdminMiddleware)
		{
			questionOpts.POST("/create", controllers.CreateQuestionOptHandler)
			questionOpts.PUT("/:questionOptID", controllers.UpdateQuestionOptHandler)
			questionOpts.DELETE("/:questionOptID", controllers.DeleteQuestionOptHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		questionOpts.GET("/", controllers.GetAllQuestionOptsHandler)
		questionOpts.GET("/:questionOptID", controllers.GetQuestionOptHandler)
	}
}
