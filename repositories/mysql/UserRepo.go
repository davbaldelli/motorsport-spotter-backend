package mysql

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"motorsportspotter.backend/models"
	"motorsportspotter.backend/repositories/entities"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u UserRepositoryImpl) Login(user models.User) (models.User, error) {
	var dbUser entities.User
	res := u.Db.Find(&dbUser, "username = ? AND password = SHA2(CONCAT(?, salt),?)", user.Username, user.Password, 224)
	if res.Error != nil {
		return models.User{}, res.Error
	}
	if res.RowsAffected == 0 {
		return models.User{}, errors.New("username or password not valid")
	}
	return models.User{Username: dbUser.Username, Role: models.Role(dbUser.Role)}, nil
}

func (u UserRepositoryImpl) SignIn(user models.User) (models.User, error) {
	salt := randStringRunes(30)
	dbUser := map[string]interface{}{
		"Username": user.Username,
		"Password": clause.Expr{SQL: "SHA2(CONCAT(?, ?),?)", Vars: []interface{}{user.Password, salt, 224}},
		"Role":     string(user.Role),
		"Salt":     salt,
	}
	if res := u.Db.Model(entities.User{}).Create(&dbUser); res.Error != nil {
		return models.User{}, res.Error
	}
	return models.User{Username: user.Username, Role: user.Role}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
