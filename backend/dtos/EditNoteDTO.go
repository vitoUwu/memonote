package dtos

type EditNoteContent struct {
	Content *string `json:"content,omitempty"`
	Title   *string `json:"title,omitempty"`
}
