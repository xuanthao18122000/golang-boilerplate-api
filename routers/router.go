package router

import (
	controller "golang-boilerplate-api/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(postsController *controller.PostsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome home")
	})
	baseRouter := router.Group("/api")
	postsRouter := baseRouter.Group("/posts")
	postsRouter.GET("", postsController.GetListPost)
	postsRouter.GET("/:postId", postsController.GetPostByID)
	postsRouter.POST("", postsController.CreatePost)
	postsRouter.PUT("/:postId", postsController.UpdatePost)
	postsRouter.DELETE("/:postId", postsController.DeletePost)

	return router
}
