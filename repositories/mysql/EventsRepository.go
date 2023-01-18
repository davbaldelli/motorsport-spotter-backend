package mysql

import (
	"errors"
	"gorm.io/gorm"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type EventsRepositoryImpl struct {
	Db *gorm.DB
}

func (e EventsRepositoryImpl) GetAllEvents() ([]models.Event, error) {
	var dbEvents []entities.DbEventView
	if result := e.Db.Order("start_date").Preload("Sessions").Find(&dbEvents); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbEventList(dbEvents).ConvertAll(), nil
}

func (e EventsRepositoryImpl) GetIncomingEvents() ([]models.Event, error) {
	var dbEvents []entities.DbEventView
	if result := e.Db.Order("start_date").Where("start_date >= CURDATE()").Preload("Sessions").Find(&dbEvents); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbEventList(dbEvents).ConvertAll(), nil
}

func (e EventsRepositoryImpl) InsertEvent(event models.Event) error {
	dbEvent := entities.DbEventFromModel(event)
	if res := e.Db.Create(&dbEvent); res.Error != nil {
		return res.Error
	}
	return nil
}

func (e EventsRepositoryImpl) UpdateEvent(event models.Event) error {
	dbEvent := entities.DbEventFromModel(event)
	if res := e.Db.Updates(&dbEvent); res.Error != nil {
		return res.Error
	}
	return nil
}
