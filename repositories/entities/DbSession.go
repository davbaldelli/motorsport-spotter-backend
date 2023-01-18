package entities

import (
	"motorsportspotter.backend/models"
)

type DbSession struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	EventId      int
	Date         string
	Time         string
	DurationMin  int
	DurationLaps int
}

type DbSessionList []DbSession

func (DbSession) TableName() string {
	return "sessions"
}

func (s DbSession) ToModel() models.Session {
	return models.Session{
		Id:           s.Id,
		Name:         s.Name,
		EventId:      s.EventId,
		Date:         s.Date,
		Time:         s.Time,
		DurationMin:  s.DurationMin,
		DurationLaps: s.DurationLaps,
	}
}

func (l DbSessionList) ConvertAll() []models.Session {
	var sessions []models.Session
	for _, sessionDb := range l {
		sessions = append(sessions, sessionDb.ToModel())
	}
	return sessions
}
