package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func QuestionEssayRoutes(r *gin.RouterGroup, questionEssayController *controllers.QuestionEssayController, authService *services.AuthService) {
	questionEssays := r.Group("/questionEssays")
	{
		// Authenticated routes
		questionEssays.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		questionEssays.Use(middlewares.AdminMiddleware)
		{
			questionEssays.POST("/create", controllers.CreateQuestionEssayHandler)
			questionEssays.PUT("/:questionEssayID", controllers.UpdateQuestionEssayHandler)
			questionEssays.DELETE("/:questionEssayID", controllers.DeleteQuestionEssayHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		questionEssays.GET("/", controllers.GetAllQuestionEssaysHandler)
		questionEssays.GET("/:questionEssayID", controllers.GetQuestionEssayHandler)
	}
}
