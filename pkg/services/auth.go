package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/OlenEnkeli/go_todo_pet/configs"
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo repositories.Authorization
}

func NewAuthService(repo repositories.Authorization) *AuthService {
	return &AuthService{repo}
}

func (srv *AuthService) CreateUser(user dto.User) (dto.User, error) {
	user.Password = srv.generatePasswordHash(user.Password)
	return srv.repo.CreateUser(user)
}

func (srv *AuthService) Login(origin dto.UserLogin) (string, error) {
	user, err := srv.repo.GetUserByLogin(origin.Login)
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("unable to login: no user with login %s", origin.Login),
		)
	}

	if user.Password != srv.generatePasswordHash(origin.Password) {
		return "", errors.New("unable to login: wrong password")
	}

	return srv.generateToken(user.Id)
}

func (srv *AuthService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(configs.Config.Auth.Salt)))
}

func (srv *AuthService) checkPassword(password string, hash string) bool {
	hashed := srv.generatePasswordHash(password)
	return hashed == hash
}

func (srv *AuthService) generateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": string(id),
		},
	)

	result, err := token.SignedString(
		[]byte(configs.Config.Auth.JWTSecretKey),
	)
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("unable to login: can`t make JWT token: %s", err),
		)
	}

	return result, nil
}
