package models

type Comment struct {
	ID          string      `json:"id"`
	Message     string      `json:"message"`
	Owner       UserPreview `json:"owner"`
	Post        string      `json:"post"`
	PublishDate string      `json:"publishDate"`
}

type CommentResponse struct {
	Data  []Comment `json:"data"`
	Total int       `json:"total"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
}
