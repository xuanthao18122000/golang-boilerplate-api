package controller

import (
	response "golang-boilerplate-api/data/response"
	"golang-boilerplate-api/models"
	service "golang-boilerplate-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostsController struct {
	postsService service.PostsService
	DB           *gorm.DB
}

type Post struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func NewPostsController(service service.PostsService) *PostsController {
	return &PostsController{
		postsService: service,
	}
}

func (controller *PostsController) isIDEmpty(id string, c *gin.Context) bool {
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID Not Found!",
		})
		return true
	}
	return false
}

func (controller *PostsController) GetListPost(c *gin.Context) {
	postsResponse := controller.postsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   postsResponse,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func idNotFoundResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, response.Response{
		Code:   http.StatusNotFound,
		Status: "ID Not Found",
		Data:   nil,
	})
}

func (controller *PostsController) GetPostByID(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		idNotFoundResponse(ctx)
		return
	}

	postResponse := controller.postsService.FindById(id)

	if postResponse.Id == 0 {
		idNotFoundResponse(ctx)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   postResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PostsController) CreatePost(c *gin.Context) {
	var post Post // Đảm bảo bạn đã định nghĩa struct Post trong mã của bạn

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Request Failed",
		})
		return
	}

	if err := controller.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Create Post!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Post Successful!",
	})
}

func (controller *PostsController) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	updatedPost := Post{}

	if controller.isIDEmpty(id, ctx) {
		return
	}

	if err := ctx.ShouldBindJSON(&updatedPost); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Failed to parse request",
		})
		return
	}

	if err := controller.DB.Model(&Post{}).Where("id = ?", id).Updates(&updatedPost).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update Post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update Post Successful!",
	})
}

func (controller *PostsController) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	postModels := models.Posts{}

	if controller.isIDEmpty(id, ctx) {
		return
	}

	result := controller.DB.Delete(&postModels, id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Delete Post!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Post Successful!",
	})
}
