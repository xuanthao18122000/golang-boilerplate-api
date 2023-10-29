package service

import (
	repository "golang-boilerplate-api/repositories"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	PostsRepository repository.PostsRepository
	Validate        *validator.Validate
}

func NewAuthServiceImpl(postRepository repository.PostsRepository, validate *validator.Validate) PostsService {
	return &PostServiceImpl{
		PostsRepository: postRepository,
		Validate:        validate,
	}
}
