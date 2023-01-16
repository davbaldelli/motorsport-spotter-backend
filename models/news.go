package models

type News struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Image    string `json:"image"`
}
