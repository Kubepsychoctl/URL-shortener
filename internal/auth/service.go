package auth

import (
	"app/url-shorter/internal/user"
	"errors"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewUserService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email string, password string, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserAlreadyExists)
	}

	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil

}
