package routes

import (
	"assignment-3/controllers"
	"assignment-3/middleware"
	"assignment-3/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authService services.AuthService, authController controllers.UserController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/sign-in", authController.SignIn)
		authRoutes.POST("/sign-up", authController.SignUp)
		authRoutes.GET("/me", middleware.Authentication(authService), authController.Profile)
	}
}
