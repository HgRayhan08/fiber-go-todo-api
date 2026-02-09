package dto

type CategoryData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
