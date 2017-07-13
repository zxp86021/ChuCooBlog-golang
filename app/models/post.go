package models

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Author Author `json:"author"`
	Tags []string `json:"tags"`
}