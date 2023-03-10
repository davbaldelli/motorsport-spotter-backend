package entities

import (
	"motorsportspotter.backend/models"
)

type DbEventView struct {
	DbEvent
	ChampionshipName       string
	ChampionshipPrettyName string
	ChampionshipYear       int
	ChampionshipLogo       string
	TrackName              string
	TrackCommonName        string
	Sessions               []DbSession `gorm:"foreignKey:EventId"`
}

type DbEvent struct {
	Id             int `gorm:"primaryKey"`
	Name           string
	Description    string
	TrackId        int
	ChampionshipId int
	StartDate      string
	EndDate        string
	Image          string
	Sessions       []DbSession `gorm:"foreignKey:EventId"`
}

type DbEventList []DbEventView

func (DbEvent) TableName() string {
	return "events"
}

func (DbEventView) TableName() string {
	return "events_view"
}

func (e DbEventView) ToModel() models.Event {
	return models.Event{
		Id:             e.Id,
		Name:           e.Name,
		Description:    e.Description,
		TrackId:        e.TrackId,
		ChampionshipId: e.ChampionshipId,
		StartDate:      e.StartDate,
		EndDate:        e.EndDate,
		Image:          e.Image,
		Championship: models.Championship{
			Name:       e.ChampionshipName,
			PrettyName: e.ChampionshipPrettyName,
			Year:       e.ChampionshipYear,
			Logo:       e.ChampionshipLogo,
		},
		Track: models.Track{
			Name:       e.TrackName,
			CommonName: e.TrackCommonName,
		},
		Sessions: DbSessionList(e.Sessions).ConvertAll(),
	}
}

func DbEventFromModel(event models.Event) DbEvent {
	return DbEvent{
		Id:             event.Id,
		Name:           event.Name,
		Description:    event.Description,
		TrackId:        event.TrackId,
		ChampionshipId: event.ChampionshipId,
		StartDate:      event.StartDate,
		EndDate:        event.EndDate,
		Image:          event.Image,
	}
}

func (l DbEventList) ConvertAll() []models.Event {
	var events []models.Event
	for _, eventDb := range l {
		events = append(events, eventDb.ToModel())
	}
	return events
}
