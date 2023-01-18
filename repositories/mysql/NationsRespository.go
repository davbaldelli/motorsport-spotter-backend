package mysql

import (
	"errors"
	"gorm.io/gorm"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type NationsRepositoryImpl struct {
	Db *gorm.DB
}

func (n NationsRepositoryImpl) GetAllNations() ([]models.Nation, error) {
	var dbNations []entities.DbNation
	if result := n.Db.Order("nation_name").Find(&dbNations); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbNationList(dbNations).ConvertAll(), nil
}
