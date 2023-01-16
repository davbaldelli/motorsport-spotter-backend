package mysql

import (
	"errors"
	"gorm.io/gorm"
	view "motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type NewsRepositoryImpl struct {
	Db *gorm.DB
}

func (r NewsRepositoryImpl) GetAllNews() ([]view.News, error) {
	var dbNews []entities.NewsDb
	if result := r.Db.Find(&dbNews); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbNewsList(dbNews).ConvertAll(), nil
}
