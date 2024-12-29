package services

import (
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/repositories"
	"math"
	"time"
)

// RaceService is the interface that wraps the basic methods to interact with the database
type RaceService interface {
	CreateRace(runnerId int, distance float64, duration float64, typeOfRunning string) (*entities.Race, error)
	GetRaces() ([]entities.Race, error)
	FindById(id int) (*entities.Race, error)
	DeleteRace(id int) error
	CalculatePace(distance, duration float64) float64
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
	pace := s.CalculatePace(distance, duration)

	race := &entities.Race{
		RunnerID:      runnerId,
		Distance:      distance,
		Duration:      duration,
		Pace:          pace,
		TypeOfRunning: typeOfRunning,
		Date:          time.Now(),
	}

	err := s.repo.Create(race)
	if err != nil {
		return nil, err
	}
	return race, nil
}

func (s *raceService) GetRaces() ([]entities.Race, error) {
	races, err := s.repo.GetRaces()
	if err != nil {
		return nil, err
	}
	return races, nil
}

func (s *raceService) FindById(id int) (*entities.Race, error) {
	return s.repo.FindById(id)
}

func (s *raceService) DeleteRace(id int) error {
	return s.repo.Delete(id)
}

func (s *raceService) CalculatePace(distance, duration float64) float64 {
	// race, err := s.repo.FindById(raceID)
	// if err != nil {
	// 	return 0, err
	// }
	pace := ((duration * 60) / distance)
	pace = math.Round(pace*100) / 100
	return pace
}
