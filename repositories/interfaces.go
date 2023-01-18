package repositories

import (
	"motorsportspotter.backend/models"
)

type ChampionshipsRepository interface {
	GetAllChampionships() ([]models.Championship, error)
	InsertChampionship(championship models.Championship) error
	UpdateChampionship(championship models.Championship) error
}

type TracksRepository interface {
	GetAllTracks() ([]models.Track, error)
	InsertTrack(models.Track) error
	UpdateTrack(models.Track) error
}

type EventsRepository interface {
	GetAllEvents() ([]models.Event, error)
	GetIncomingEvents() ([]models.Event, error)
	InsertEvent(event models.Event) error
	UpdateEvent(event models.Event) error
}

type SessionsRepository interface {
	GetAllSessions() ([]models.Session, error)
	InsertSession(session models.Session) error
	UpdateSession(session models.Session) error
}

type NationsRepository interface {
	GetAllNations() ([]models.Nation, error)
}

type NewsRepository interface {
	GetAllNews() ([]models.News, error)
}

type UserRepository interface {
	Login(user models.User) (models.User, error)
	SignIn(user models.User) (models.User, error)
}
