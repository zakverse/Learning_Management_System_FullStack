package domain

import "time"

type Chapter struct {	
	ID        uint      `gorm:"primaryKey"`
	CourseID  uint      `gorm:"not null;index"`
	Title     string    `gorm:"size:200;not null"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
}

type ChapterRepository interface {
	Create(chapter *Chapter) error
	FindByCourseID(courseID uint) ([]Chapter, error)
	FindByID(id uint) (*Chapter, error)
	FindAll() ([]Chapter, error)
}
