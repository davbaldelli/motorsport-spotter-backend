package entities

import "motorsportspotter.backend/models"

type DbNation struct {
	NationCode string `gorm:"primaryKey"`
	NationName string
}

type DbNationList []DbNation

func (DbNation) TableName() string {
	return "nations"
}

func (n DbNation) ToModel() models.Nation {
	return models.Nation{
		Code: n.NationCode,
		Name: n.NationName,
	}
}

func DbNationFromModel(nation models.Nation) DbNation {
	return DbNation{
		NationCode: nation.Code,
		NationName: nation.Name,
	}
}

func (l DbNationList) ConvertAll() []models.Nation {
	var nations []models.Nation
	for _, dbNation := range l {
		nations = append(nations, dbNation.ToModel())
	}
	return nations
}
