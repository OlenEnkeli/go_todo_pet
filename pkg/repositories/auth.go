package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories/models"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func (repo *AuthDB) GetCurrentUser(userId int) (dto.User, error) {
	var user *models.User

	result := repo.db.
		Model(&user).
		Where("id = ?", userId).
		First(&user)

	return user.ToDTO(), result.Error
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

	err := repo.db.Create(&newUser)

	return newUser.ToDTO(), err.Error
}

func (repo *AuthDB) GetUserByLogin(login string) (dto.User, error) {
	var user *models.User

	result := repo.db.
		Model(&models.User{}).
		Where("login = ?", login).
		First(&user)

	return user.ToDTO(), result.Error
}
