package models

type Post struct {
	ID          string      `json:"id"`
	Text        string      `json:"text"`
	Likes       int         `json:"likes"`
	Tags        []string    `json:"tags"`
	PublishDate string      `json:"publishDate"`
	User        UserPreview `json:"owner"`
}

type PostResponse struct {
	Data  []Post `json:"data"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
