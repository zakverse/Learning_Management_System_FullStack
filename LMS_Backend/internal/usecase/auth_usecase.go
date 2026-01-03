package usecase

import (
	"LMS_Backend/internal/domain"
	"LMS_Backend/internal/infrastructure"
	"errors"
)

type AuthUsecase interface {
	Register(name, email, password, role string) error
	Login(email, password string) (*domain.User, string, error)
}

type authUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(userRepo domain.UserRepository) AuthUsecase {
	return &authUsecase{userRepo}
}

func (u *authUsecase) Register(name, email, password, role string) error {
	hashed, err := infrastructure.HashPassword(password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     role,
	}

	return u.userRepo.Create(user)
}

func (u *authUsecase) Login(email, password string) (*domain.User, string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("email tidak ditemukan")
	}

	if !infrastructure.CheckPassword(user.Password, password) {
		return nil, "", errors.New("password salah")
	}

	token, err := infrastructure.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}
