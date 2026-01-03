package domain

import "time"

type Course struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:200;not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}


type CourseRepository interface {
	Create(course *Course) error
	FindAll() ([]Course, error)
	FindById(id uint) (*Course, error)
}

