package mysql

import (
	"errors"
	"gorm.io/gorm"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type SessionRepositoryImpl struct {
	Db *gorm.DB
}

func (s SessionRepositoryImpl) GetAllSessions() ([]models.Session, error) {
	var dbSessions []entities.SessionDb
	if result := s.Db.Find(&dbSessions); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbSessionList(dbSessions).ConvertAll(), nil
}

func (s SessionRepositoryImpl) InsertSession(session models.Session) error {
	if res := s.Db.Create(&session); res.Error != nil {
		return res.Error
	}
	return nil
}

func (s SessionRepositoryImpl) UpdateSession(session models.Session) error {
	if res := s.Db.Updates(&session); res.Error != nil {
		return res.Error
	}
	return nil
}
