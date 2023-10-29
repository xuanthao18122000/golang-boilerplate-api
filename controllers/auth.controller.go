package controller

import (
	service "golang-boilerplate-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	postsService service.PostsService
	DB           *gorm.DB
}

func NewAuthController(service service.PostsService) *AuthController {
	return &AuthController{
		postsService: service,
	}
}

func (controller *AuthController) Login(c *gin.Context) {

}

func (controller *AuthController) Register(c *gin.Context) {

}
