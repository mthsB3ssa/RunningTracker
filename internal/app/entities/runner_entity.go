package entities

import (
	"time"
)

type Runner struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Races     []Race    `gorm:"foreignKey:RunnerID"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
