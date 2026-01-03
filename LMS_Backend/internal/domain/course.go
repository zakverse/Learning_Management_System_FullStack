package domain

import "time"

type Course struct {
	ID          uint		`gorm:"primaryKey"`
	Title       string		`gorm:"size:200;not null"`
	Description string		`gorm:"type:text"`
	CreatedBy   uint		`gorm:"not null"`
	CreatedAt   time.Time
}

type CourseRepository interface {
	Create(course *Course) error
	FindAll() ([]Course, error)
	FindById(id uint) (*Course, error)
}

