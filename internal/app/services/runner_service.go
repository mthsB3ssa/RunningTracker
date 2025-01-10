package services

import (
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/repositories"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Define o contrato para os serviços que lidam com a lógica do negócio
type RunnerService interface {
	CreateRunner(name string, age int, email string) (*entities.Runner, error)
	GetUsers() ([]entities.Runner, error)
	FindById(id int) (*entities.Runner, error)
	UpdateRunner(runner *entities.Runner) (*entities.Runner, error)
	DeleteRunner(id int) error
	GetAllRunsByUser(id int) ([]entities.Runner, error)
}

// Implementação da interface
type runnerService struct {
	repo repositories.RunnerRepository
}

// Cria uma instância de runnerService
func NewRunnerService(repo repositories.RunnerRepository) RunnerService {
	return &runnerService{repo: repo}
}

// Cria uma nova instância de Runner, define os dados e chama o repo para salvar no banco
func (s *runnerService) CreateRunner(name string, age int, email string) (*entities.Runner, error) {
	runner := &entities.Runner{
		Name:      name,
		Age:       age,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Create(runner)
	if err != nil {
		return nil, err
	}
	return runner, nil
}

func (s *runnerService) GetUsers() ([]entities.Runner, error) {
	runners, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return runners, nil
}

func (s *runnerService) FindById(id int) (*entities.Runner, error) {
	return s.repo.FindById(id)
}

func (s *runnerService) UpdateRunner(runner *entities.Runner) (*entities.Runner, error) {
	existingRunner, err := s.repo.FindById(runner.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Runner nor found")
		}
		return nil, err
	}

	if runner.Name != "" {
		existingRunner.Name = runner.Name
	}

	if runner.Age != 0 {
		existingRunner.Age = runner.Age
	}

	existingRunner.Email = runner.Email

	existingRunner.UpdatedAt = time.Now()

	err = s.repo.Update(existingRunner)
	if err != nil {
		return nil, err
	}
	return existingRunner, nil
}

func (s *runnerService) DeleteRunner(id int) error {
	return s.repo.Delete(id)
}

func (s *runnerService) GetAllRunsByUser(id int) ([]entities.Runner, error) {
	runners, err := s.repo.GetAllRunsByUser(id)
	if err != nil {
		return nil, err
	}
	return runners, nil
}
