package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func AnswerOptRoutes(r *gin.RouterGroup, answerOptController *controllers.AnswerOptController, authService *services.AuthService) {
	answerOpts := r.Group("/answerOpts")
	{
		// Authenticated routes
		answerOpts.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		answerOpts.Use(middlewares.AdminMiddleware)
		{
			answerOpts.POST("/create", controllers.CreateAnswerOptHandler)
			answerOpts.PUT("/:answerOptID", controllers.UpdateAnswerOptHandler)
			answerOpts.DELETE("/:answerOptID", controllers.DeleteAnswerOptHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		answerOpts.GET("/:answerOptID", controllers.GetAnswerOptHandler)
	}
}
