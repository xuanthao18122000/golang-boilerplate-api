// main_router.go
package router

import (
	controller "golang-boilerplate-api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(postsController *controller.PostsController, authController *controller.AuthController) *gin.Engine {
	router := gin.Default()

	// Add Swagger documentation route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome home")
	})

	baseRouter := router.Group("/api")

	// Import and set up the post routes
	SetupPostRoutes(baseRouter, postsController)
	SetupAuthRoutes(baseRouter, authController)

	return router
}
