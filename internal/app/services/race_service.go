package services

import (
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/repositories"
	"time"

	"golang.org/x/text/date"
)

type RaceService interface {
	CreateRace(runnerId int, distance float64, duration float64, typeOfRunning string, date time) (*entities.Race, error)
	FindById (id int) error
	DeleteRace(id int) error
}

type raceService struct {
	repo repositories.RaceRepository
}

func NewRaceService(repo repositories.RaceRepository) RaceService {
	return &raceService{repo: repo}
}

func (s *raceService) CreateRace(runnerId int, distance float64, duration *float64, typeOfRunning string, date time.Time)(*entities.Race, error){
	race := &entities.Race{
		RunnerID: runnerId,
		Distance: distance,
		Duration: duration,
		TypeOfRunning: typeOfRunning,
		Date: date,
	}

	err := s.repo.Create(race)
	if err != nil {
		return nil, err
	}
	return race, nil
}

func (s *raceService) FindById(id int) error{
	return s.repo.FindById(id)
}

func (s *raceService) DeleteRace(id int) error {
	return s.repo.Delete(id)
}
