package domain

import "time"

type User struct {
	ID        uint    		`gorm:"primaryKey"`
	Name      string  		`gorm:"size:100;not null"`
	Email     string  		`gorm:"size:100;unique;not null"`
	Password  string  		`gorm:"not null"`
	Role      string  		`gorm:"type:enum('admin','dosen','mahasiswa');default:'mahasiswa'"`
	CreatedAt time.Time 	`gorm:"autoCreateTime"`
}

type UserRepository interface {
	Create(user User) error
	FindByEmail(email string) (User, error)
}
