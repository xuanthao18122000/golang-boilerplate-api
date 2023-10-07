package repository

import "golang-boilerplate-api/models"

type PostsRepository interface {
	Save(posts models.Posts)
	Update(posts models.Posts)
	Delete(postsId int)
	FindById(postsId int) (posts models.Posts, err error)
	FindAll() []models.Posts
}
