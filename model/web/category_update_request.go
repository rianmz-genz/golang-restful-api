package web

type CategoryUpdateRequest struct {
	Id int `validate:"required"`
	Name string `validate:"required,min=5" json:"name"`
}