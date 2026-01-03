package repository

import (
	"LMS_Backend/internal/domain"
	"gorm.io/gorm"
)

type chapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) domain.ChapterRepository {
	return &chapterRepository{db}
}

func (r *chapterRepository) Create(chapter *domain.Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *chapterRepository) FindByCourseID(courseID uint) ([]domain.Chapter, error) {
	var chapters []domain.Chapter
	err := r.db.Where("course_id = ?", courseID).Order("created_at desc").Find(&chapters).Error
	return chapters, err
}

func (r *chapterRepository) FindByID(id uint) (*domain.Chapter, error) {
	var chapter domain.Chapter
	err := r.db.Where("id = ?", id).Find(&chapter).Error
	return &chapter, err
}

func (r *chapterRepository) FindAll() ([]domain.Chapter, error) {
	var chapters []domain.Chapter
	err := r.db.Find(&chapters).Error
	return chapters, err
}
