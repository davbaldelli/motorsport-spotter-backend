package entities

import (
	"motorsportspotter.backend/models"
)

type EventView struct {
	EventDb
	ChampionshipName string
	ChampionshipYear int
	TrackName        string
	Sessions         []SessionDb `gorm:"foreignKey:EventId"`
}

type EventDb struct {
	Id             int `gorm:"primaryKey"`
	Name           string
	TrackId        int
	ChampionshipId int
	StartDate      string
	EndDate        string
	Image          string
	Sessions       []SessionDb `gorm:"foreignKey:EventId"`
}

type DbEventList []EventView

func (EventDb) TableName() string {
	return "events"
}

func (EventView) TableName() string {
	return "events_view"
}

func (e EventView) ToEntity() models.Event {
	return models.Event{
		Id:             e.Id,
		Name:           e.Name,
		TrackId:        e.TrackId,
		ChampionshipId: e.ChampionshipId,
		StartDate:      e.StartDate,
		EndDate:        e.EndDate,
		Image:          e.Image,
		Championship: models.Championship{
			Name: e.ChampionshipName,
			Year: e.ChampionshipYear,
		},
		Track: models.Track{
			Name: e.TrackName,
		},
		Sessions: DbSessionList(e.Sessions).ConvertAll(),
	}
}

func (l DbEventList) ConvertAll() []models.Event {
	var events []models.Event
	for _, eventDb := range l {
		events = append(events, eventDb.ToEntity())
	}
	return events
}
