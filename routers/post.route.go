package router

import (
	controller "golang-boilerplate-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(baseRouter *gin.RouterGroup, postsController *controller.PostsController) {
	postsRouter := baseRouter.Group("/posts")
	postsRouter.GET("", postsController.GetListPost)
	postsRouter.GET("/:postId", postsController.GetPostByID)
	postsRouter.POST("", postsController.CreatePost)
	postsRouter.PUT("/:postId", postsController.UpdatePost)
	postsRouter.DELETE("/:postId", postsController.DeletePost)
}
