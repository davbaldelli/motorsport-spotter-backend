package entities

import (
	"motorsportspotter.backend/models"
)

type TrackDB struct {
	Id           int
	Name         string
	Image        string
	Logo         string
	LocationName string
	NationCode   string
}

type DbTrackList []TrackDB

func (TrackDB) TableName() string {
	return "tracks"
}

func (t TrackDB) ToEntity() models.Track {
	return models.Track{
		Id:           t.Id,
		Name:         t.Name,
		Image:        t.Image,
		Logo:         t.Logo,
		LocationName: t.LocationName,
		NationCode:   t.NationCode,
	}
}

func (l DbTrackList) ConvertAll() []models.Track {
	var tracks []models.Track
	for _, trackDb := range l {
		tracks = append(tracks, trackDb.ToEntity())
	}
	return tracks
}
