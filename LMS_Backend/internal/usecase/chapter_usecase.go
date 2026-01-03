package usecase

import (
	"LMS_Backend/internal/domain"
	"errors"
)

type ChapterUsecase interface {
	Create(courseID uint, title, content string) error
	GetByCourse(courseID uint) ([]domain.Chapter, error)
	GetByID(id uint) (*domain.Chapter, error)
}

type chapterusecase struct{
	repo domain.ChapterRepository

}

func NewChapterUsecase(repo domain.ChapterRepository) ChapterUsecase {
	return &chapterusecase{repo}
}

func (u *chapterusecase) Create(courseID uint, title, content string) error {
	
	if title == "" {
		return errors.New("title wajib diisi")
	}
	
	chapter := domain.Chapter{
		CourseID: courseID,
		Title:    title,
		Content:  content,
	}
	return u.repo.Create(&chapter)
}

func (u *chapterusecase) GetByCourse(courseID uint) ([]domain.Chapter, error) {
	return u.repo.FindByCourseID(courseID)
}

func (u *chapterusecase) GetByID(id uint) (*domain.Chapter, error) {
	return u.repo.FindByID(id)
}


