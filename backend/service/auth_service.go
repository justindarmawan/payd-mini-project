package service

import (
	"errors"

	"github.com/justindarmawan/payd-mini-project/backend/model"
	"github.com/justindarmawan/payd-mini-project/backend/repository"
	"github.com/justindarmawan/payd-mini-project/backend/utils"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) Login(username, password string) (token string, user *model.User, err error) {
	user, err = s.UserRepo.GetByUsername(username)
	if err != nil {
		return "", user, errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", user, errors.New("invalid username or password")
	}

	token, err = utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", user, err
	}

	return token, user, nil
}
