package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(r *gin.RouterGroup, courseController *controllers.CourseController, authService *services.AuthService) {
	courses := r.Group("/courses")
	{
		// Authenticated routes
		courses.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		courses.Use(middlewares.AdminMiddleware)
		{
			courses.POST("/create", controllers.CreateCourseHandler)
			courses.PUT("/:courseID", controllers.UpdateCourseHandler)
			courses.DELETE("/:courseID", controllers.DeleteCourseHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		courses.GET("/", controllers.GetAllCoursesHandler)
		courses.GET("/:courseID", controllers.GetCourseByIDHandler)
	}
}
