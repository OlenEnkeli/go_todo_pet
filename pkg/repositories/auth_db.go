package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories/models"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB {
	return &AuthDB{db}
}

func (repo *AuthDB) CreateUser(user dto.User) (dto.User, error) {
	newUser := models.User{
		Login:    user.Login,
		Password: user.Password,
		Name:     user.Name,
	}

	result := repo.db.Create(&newUser)

	return newUser.ToDTO(), result.Error
}

func (repo *AuthDB) GetUserByLogin(login string) (dto.User, error) {
	var user *models.User

	repo.db.
		Model(&models.User{}).
		Where("login = ?", login).Scan(&user)

	return user.ToDTO(), nil
}
