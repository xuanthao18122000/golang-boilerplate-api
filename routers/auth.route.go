package router

import (
	controller "golang-boilerplate-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(baseRouter *gin.RouterGroup, authController *controller.AuthController) {
	postsRouter := baseRouter.Group("/auth")
	postsRouter.GET("/login", authController.Login)
	postsRouter.GET("/register", authController.Register)
}
