package routes

import (
	"shopping-website/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup, authController *controllers.AuthController) {
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	// authRoutes.PUT("/update", authController.Update) // Assuming Update method exists in AuthController
	// authRoutes.POST("/logout", authController.Logout) // Assuming Logout method exists in AuthController

}
