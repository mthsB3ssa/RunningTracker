package entities

import "time"

type Race struct {
	ID             uint      `json:"id"`
	Duration       time.Time `json:"duration"`
	TotalKm        float32   `json:"total_km"`
	TypeOfTraining string    `json:"type_of_training"`
}
