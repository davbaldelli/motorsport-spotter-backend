package mysql

import (
	"errors"
	"gorm.io/gorm"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type ChampionshipRepoImpl struct {
	Db *gorm.DB
}

func (r ChampionshipRepoImpl) GetAllChampionships() ([]models.Championship, error) {
	var dbChamps []entities.ChampionshipDb
	if result := r.Db.Find(&dbChamps); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return entities.DbChampionshipList(dbChamps).ConvertAll(), nil
}

func (r ChampionshipRepoImpl) InsertChampionship(championship models.Championship) error {
	if res := r.Db.Create(&championship); res.Error != nil {
		return res.Error
	}
	return nil
}

func (r ChampionshipRepoImpl) UpdateChampionship(championship models.Championship) error {
	if res := r.Db.Updates(&championship); res.Error != nil {
		return res.Error
	}
	return nil
}
