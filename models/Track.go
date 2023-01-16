package models

type Track struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Logo         string `json:"logo"`
	LocationName string `json:"locationName"`
	NationCode   string `json:"nationCode"`
}
