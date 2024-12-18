package services

import (
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/repositories"
	"time"
)

// RaceService is the interface that wraps the basic methods to interact with the database
type RaceService interface {
	CreateRace(runnerId int, distance float64, duration float64, typeOfRunning string) (*entities.Race, error)
	FindById(id int) (*entities.Race, error)
	DeleteRace(id int) error
}

// raceService is the struct that implements the RaceService interface
type raceService struct {
	repo repositories.RaceRepository
}

// NewRaceService creates a new instance of the raceService struct
func NewRaceService(repo repositories.RaceRepository) RaceService {
	return &raceService{repo: repo}
}

func (s *raceService) CreateRace(runnerId int, distance float64, duration float64, typeOfRunning string) (*entities.Race, error) {
	race := &entities.Race{
		RunnerID:      runnerId,
		Distance:      distance,
		Duration:      duration,
		TypeOfRunning: typeOfRunning,
		Date:          time.Now(),
	}

	err := s.repo.Create(race)
	if err != nil {
		return nil, err
	}
	return race, nil
}

func (s *raceService) FindById(id int) (*entities.Race, error) {
	return s.repo.FindById(id)
}

func (s *raceService) DeleteRace(id int) error {
	return s.repo.Delete(id)
}
