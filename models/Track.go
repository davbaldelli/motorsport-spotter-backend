package models

type Track struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CommonName   string `json:"commonName"`
	Image        string `json:"image"`
	Logo         string `json:"logo"`
	Description  string `json:"description"`
	LocationName string `json:"locationName"`
	NationCode   string `json:"nationCode"`
	Nation       Nation `json:"nation"`
}
