package service

import (
	"golang-boilerplate-api/data/request"
	"golang-boilerplate-api/data/response"
)

type PostsService interface {
	Create(Posts request.CreatePostsRequest)
	Update(Posts request.UpdatePostsRequest)
	Delete(PostsId int)
	FindById(PostsId int) response.PostsResponse
	FindAll() []response.PostsResponse
}
