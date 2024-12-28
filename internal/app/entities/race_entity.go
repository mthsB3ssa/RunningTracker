package entities

import "time"

type Race struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	RunnerID      int       `json:"-"`
	Distance      float64   `json:"distance"`
	Pace          float64   `json:"pace"`
	Duration      float64   `json:"duration"`
	TypeOfRunning string    `json:"type"`
	Date          time.Time `json:"date"`
}
