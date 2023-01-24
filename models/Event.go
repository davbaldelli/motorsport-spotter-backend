package models

type Event struct {
	Id             int          `json:"id"`
	Name           string       `json:"name"`
	TrackId        int          `json:"trackId"`
	ChampionshipId int          `json:"championshipId"`
	StartDate      string       `json:"startDate"`
	EndDate        string       `json:"endDate"`
	Image          string       `json:"image"`
	Description    string       `json:"description"`
	Championship   Championship `json:"championship"`
	Track          Track        `json:"track"`
	Sessions       []Session    `json:"sessions"`
}
