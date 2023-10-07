package repository

import (
	"errors"
	"golang-boilerplate-api/data/request"
	helper "golang-boilerplate-api/helpers"
	"golang-boilerplate-api/models"

	"gorm.io/gorm"
)

type PostsRepositoryImpl struct {
	Db *gorm.DB
}

func NewPostsREpositoryImpl(Db *gorm.DB) PostsRepository {
	return &PostsRepositoryImpl{Db: Db}
}

// Delete implements PostsRepository
func (t *PostsRepositoryImpl) Delete(postsId int) {
	var posts models.Posts
	result := t.Db.Where("id = ?", postsId).Delete(&posts)
	helper.ErrorPanic(result.Error)
}

// FindAll implements PostsRepository
func (t *PostsRepositoryImpl) FindAll() []models.Posts {
	var posts []models.Posts
	result := t.Db.Find(&posts)
	helper.ErrorPanic(result.Error)
	return posts
}

// FindById implements PostsRepository
func (t *PostsRepositoryImpl) FindById(postsId int) (posts models.Posts, err error) {
	var post models.Posts
	result := t.Db.Find(&post, postsId)
	if result != nil {
		return post, nil
	} else {
		return post, errors.New("post is not found")
	}
}

// Save implements PostsRepository
func (t *PostsRepositoryImpl) Save(posts models.Posts) {
	result := t.Db.Create(&posts)
	helper.ErrorPanic(result.Error)
}

// Update implements PostsRepository
func (t *PostsRepositoryImpl) Update(posts models.Posts) {
	var updatePost = request.UpdatePostsRequest{
		Id:   int(posts.Id),
		Name: *posts.Name,
	}
	result := t.Db.Model(&posts).Updates(updatePost)
	helper.ErrorPanic(result.Error)
}
