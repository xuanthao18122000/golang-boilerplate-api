package request

type UpdatePostsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}