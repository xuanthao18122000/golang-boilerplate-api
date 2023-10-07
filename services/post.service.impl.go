package service

import (
	"golang-boilerplate-api/data/request"
	"golang-boilerplate-api/data/response"
	helper "golang-boilerplate-api/helpers"
	"golang-boilerplate-api/models"
	repository "golang-boilerplate-api/repositories"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostsRepository repository.PostsRepository
	Validate        *validator.Validate
}

func NewPostServiceImpl(postRepository repository.PostsRepository, validate *validator.Validate) PostsService {
	return &PostServiceImpl{
		PostsRepository: postRepository,
		Validate:        validate,
	}
}

// Create implements PostsService
func (t *PostServiceImpl) Create(posts request.CreatePostsRequest) {
	err := t.Validate.Struct(posts)
	helper.ErrorPanic(err)
	postModel := models.Posts{
		Name: &posts.Name,
	}
	t.PostsRepository.Save(postModel)
}

// Delete implements PostsService
func (t *PostServiceImpl) Delete(postsId int) {
	t.PostsRepository.Delete(postsId)
}

// FindAll implements PostsService
func (t *PostServiceImpl) FindAll() []response.PostsResponse {
	result := t.PostsRepository.FindAll()

	var posts []response.PostsResponse
	for _, value := range result {
		post := response.PostsResponse{
			Id:   int(value.Id),
			Name: *value.Name,
		}
		posts = append(posts, post)
	}

	return posts
}

// FindById implements PostsService
func (t *PostServiceImpl) FindById(postsId int) response.PostsResponse {
	postData, err := t.PostsRepository.FindById(postsId)
	helper.ErrorPanic(err)

	postResponse := response.PostsResponse{
		Id:   int(postData.Id),
		Name: *postData.Name,
	}
	return postResponse
}

// Update implements PostsService
func (t *PostServiceImpl) Update(posts request.UpdatePostsRequest) {
	postData, err := t.PostsRepository.FindById(posts.Id)
	helper.ErrorPanic(err)
	postData.Name = &posts.Name
	t.PostsRepository.Update(postData)
}
