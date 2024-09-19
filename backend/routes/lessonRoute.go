package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func LessonRoutes(r *gin.RouterGroup, lessonController *controllers.LessonController, authService *services.AuthService) {
	lessons := r.Group("/lessons")
	{
		// Authenticated routes
		lessons.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		lessons.Use(middlewares.AdminMiddleware)
		{
			lessons.POST("/create", controllers.CreateLessonHandler)
			lessons.PUT("/:lessonID", controllers.UpdateLessonHandler)
			lessons.DELETE("/:lessonID", controllers.DeleteLessonHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		lessons.GET("/:lessonID", controllers.GetLessonHandler)
	}
}
