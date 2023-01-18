package mysql

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type TracksRepositoryImpl struct {
	Db *gorm.DB
}

func (t TracksRepositoryImpl) preInsertionQueries(track models.Track) error {

	dbNation := entities.DbNationFromModel(track.Nation)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	return nil
}

func (t TracksRepositoryImpl) GetAllTracks() ([]models.Track, error) {
	var dbTracks []entities.DbTrackView
	if result := t.Db.Find(&dbTracks); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbTrackList(dbTracks).ConvertAll(), nil
}

func (t TracksRepositoryImpl) InsertTrack(track models.Track) error {
	if err := t.preInsertionQueries(track); err != nil {
		return err
	}
	dbTrack := entities.DbTrackFromModel(track)
	if res := t.Db.Create(&dbTrack); res.Error != nil {
		return res.Error
	}
	return nil
}

func (t TracksRepositoryImpl) UpdateTrack(track models.Track) error {
	if err := t.preInsertionQueries(track); err != nil {
		return err
	}
	dbTrack := entities.DbTrackFromModel(track)
	if res := t.Db.Updates(&dbTrack); res.Error != nil {
		return res.Error
	}
	return nil
}
