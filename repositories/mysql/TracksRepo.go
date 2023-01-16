package mysql

import (
	"errors"
	"gorm.io/gorm"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type TracksRepositoryImpl struct {
	Db *gorm.DB
}

func (t TracksRepositoryImpl) GetAllTracks() ([]models.Track, error) {
	var dbTracks []entities.TrackDB
	if result := t.Db.Find(&dbTracks); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbTrackList(dbTracks).ConvertAll(), nil
}

func (t TracksRepositoryImpl) InsertTrack(track models.Track) error {
	if res := t.Db.Create(&track); res.Error != nil {
		return res.Error
	}
	return nil
}

func (t TracksRepositoryImpl) UpdateTrack(track models.Track) error {
	if res := t.Db.Updates(&track); res.Error != nil {
		return res.Error
	}
	return nil
}
