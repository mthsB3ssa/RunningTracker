package entities

import (
	"time"
)

type Runner struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Races     []Race    `gorm:"foreignKey:RunnerID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
