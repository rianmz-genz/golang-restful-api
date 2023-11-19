package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=5" json:"name"`
}