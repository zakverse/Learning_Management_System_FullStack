package repository

import (
	"LMS_Backend/internal/domain"

	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) domain.CourseRepository {
	return &courseRepository{db: db}
}

func (c *courseRepository) Create(course *domain.Course) error {
	return c.db.Create(&course).Error
}

func (r *courseRepository) FindAll() ([]domain.Course, error) {
	var courses []domain.Course
	err := r.db.Order("created_at desc").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindById(id uint) (*domain.Course, error) {
	var course domain.Course
	err := r.db.First(&course, id).Error
	return &course, err
}

func (r *courseRepository) Update(course *domain.Course) error {	
	return r.db.Save(course).Error
}
