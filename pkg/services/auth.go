package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/OlenEnkeli/go_todo_pet/configs"
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
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

	return srv.generateToken(user)
}

func (srv *AuthService) GetCurrentUser(userId int) (dto.User, error) {
	return srv.repo.GetCurrentUser(userId)
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

func (srv *AuthService) generateToken(user dto.User) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   strconv.Itoa(int(user.Id)),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	result, err := token.SignedString([]byte(configs.Config.Auth.JWTSecretKey))
	if err != nil {
		return "", errors.New(
			fmt.Sprintf("unable to login: can`t make JWT token: %s", err),
		)
	}

	return result, nil
}

func (srv *AuthService) ParseToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.Config.Auth.JWTSecretKey), nil
		},
	)
	if err != nil {
		return 0, errors.New("wrong JWT token: can`t parse")
	}

	if !parsedToken.Valid {
		return 0, errors.New("wrong JWT token: token isn`t valid")
	}

	id, err := parsedToken.Claims.GetSubject()

	if err != nil {
		return 0, errors.New("wrong JWT token: no sub field")
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf(err.Error())
		return 0, errors.New("wrong JWT token: id must be integer")
	}

	return parsedId, nil
}
