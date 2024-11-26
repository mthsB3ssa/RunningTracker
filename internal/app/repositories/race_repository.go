package repositories

import (
	"RunningTracker/internal/app/entities"

	"gorm.io/gorm"
)

type RaceRepository interface {
	Create(race *entities.Race) error
	Delete(race *entities.Race) error
}

type raceRepository struct {
	db *gorm.DB
}

func NewRaceRepository(db *gorm.DB) RaceRepository {
	return &raceRepository{db: db}
}

func (repo *raceRepository) Create(race *entities.Race) error {
	return repo.db.Create(race).Error
}

func (repo *raceRepository) Delete(race *entities.Race) error {
	return nil
}
