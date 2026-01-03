package usecase

import (
	"errors"

	"LMS_Backend/internal/domain"
	
)

type CourseUsecase interface {
	Create(title, description string, createdBy uint) error
	GetAll() ([]domain.Course, error)
	GetByID(id uint) (*domain.Course, error)
}

type courseUsecase struct {
	repo domain.CourseRepository
}

func NewCourseUsecase(repo domain.CourseRepository) CourseUsecase {
	return &courseUsecase{repo}
}

func (u *courseUsecase) Create(title, description string, createdBy uint) error {
	if title == "" {
		return errors.New("title wajib diisi")
	}

	course := domain.Course{
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
	}

	return u.repo.Create(&course)
}

func (u *courseUsecase) GetAll() ([]domain.Course, error) {
	return u.repo.FindAll()
}

func (u *courseUsecase) GetByID(id uint) (*domain.Course, error) {
	return u.repo.FindById(id)
}
