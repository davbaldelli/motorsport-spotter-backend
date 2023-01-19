package entities

import (
	"motorsportspotter.backend/models"
)

type DbTrack struct {
	Id           int
	Name         string
	CommonName   string
	Image        string
	Logo         string
	LocationName string
	NationCode   string
}

type DbTrackView struct {
	DbTrack
	NationName string
}

type DbTrackList []DbTrackView

func (DbTrack) TableName() string {
	return "tracks"
}

func (DbTrackView) TableName() string {
	return "tracks_view"
}

func DbTrackFromModel(track models.Track) DbTrack {
	return DbTrack{
		Id:           track.Id,
		Name:         track.Name,
		CommonName:   track.CommonName,
		Image:        track.Image,
		Logo:         track.Logo,
		LocationName: track.LocationName,
		NationCode:   track.Nation.Code,
	}
}

func (t DbTrackView) ToModel() models.Track {
	return models.Track{
		Id:           t.Id,
		Name:         t.Name,
		CommonName:   t.CommonName,
		Image:        t.Image,
		Logo:         t.Logo,
		LocationName: t.LocationName,
		Nation: models.Nation{
			Code: t.NationCode,
			Name: t.NationName,
		},
	}
}

func (l DbTrackList) ConvertAll() []models.Track {
	var tracks []models.Track
	for _, trackDb := range l {
		tracks = append(tracks, trackDb.ToModel())
	}
	return tracks
}
