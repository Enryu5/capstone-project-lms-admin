package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func ModuleRoutes(r *gin.RouterGroup, moduleController *controllers.ModuleController, authService *services.AuthService) {
	modules := r.Group("/modules")
	{
		// Authenticated routes
		modules.Use(middlewares.AuthMiddleware(authService))

		// Admin routes
		modules.Use(middlewares.AdminMiddleware)
		{
			modules.POST("/create", controllers.CreateModuleHandler)
			modules.PUT("/:moduleID", controllers.UpdateModuleHandler)
			modules.DELETE("/:moduleID", controllers.DeleteModuleHandler)
		}

		// General routes (can be accessed by anyone authenticated)
		modules.GET("/:moduleID", controllers.GetModuleHandler)
	}
}
