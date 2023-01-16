package entities

import (
	"motorsportspotter.backend/models"
)

type DbChampionshipList []ChampionshipDb

type ChampionshipDb struct {
	Id         int
	Name       string
	PrettyName string
	Year       int
	Image      string
	Logo       string
	LiveStream string
}

func (ChampionshipDb) TableName() string {
	return "championships"
}

func (c ChampionshipDb) ToEntity() models.Championship {
	return models.Championship{
		Id:         c.Id,
		Name:       c.Name,
		PrettyName: c.PrettyName,
		Year:       c.Year,
		Image:      c.Image,
		Logo:       c.Logo,
		LiveStream: c.LiveStream,
	}
}

func (a DbChampionshipList) ConvertAll() []models.Championship {
	var championships []models.Championship
	for _, champDb := range a {
		championships = append(championships, champDb.ToEntity())
	}
	return championships
}
