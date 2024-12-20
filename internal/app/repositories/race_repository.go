package repositories

import (
	"RunningTracker/internal/app/entities"

	"gorm.io/gorm"
)

// RaceRepository is the interface that wraps the basic methods to interact with the database
type RaceRepository interface {
	Create(race *entities.Race) error
	FindById(id int) (*entities.Race, error)
	Delete(id int) error
}

// raceRepository is the struct that implements the RaceRepository interface
type raceRepository struct {
	db *gorm.DB
}

// NewRaceRepository creates a new instance of the raceRepository struct
func NewRaceRepository(db *gorm.DB) RaceRepository {
	return &raceRepository{db: db}
}

func (repo *raceRepository) Create(race *entities.Race) error {
	return repo.db.Create(race).Error
}

func (repo raceRepository) FindById(id int) (*entities.Race, error) {
	var race entities.Race
	err := repo.db.First(&race, id).Error
	if err != nil {
		return nil, err
	}
	return &race, nil
}

func (repo *raceRepository) Delete(id int) error {
	return nil
}
