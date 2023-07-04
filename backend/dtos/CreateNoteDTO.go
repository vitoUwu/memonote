package dtos

type CreateNote struct {
	Content string `json:"content" validate:"required"`
	Title   string `json:"title" validate:"required"`
}
