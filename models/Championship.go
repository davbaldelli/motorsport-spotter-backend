package models

type Championship struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PrettyName string `json:"prettyName"`
	Year       int    `json:"year"`
	Image      string `json:"image"`
	Logo       string `json:"logo"`
	LiveStream string `json:"liveStream"`
}
