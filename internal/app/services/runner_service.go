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
	CreateRunner(name string, age int) (*entities.Runner, error)
	GetUsers() ([]entities.Runner, error)
	FindById(id int) (*entities.Runner, error)
	UpdateRunner(id int, name string, age int) (*entities.Runner, error)
	DeleteRunner(id int) error
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
func (s *runnerService) CreateRunner(name string, age int) (*entities.Runner, error) {
	runner := &entities.Runner{
		Name:      name,
		Age:       age,
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
	result, err := s.repo.GetUsers()

	var runners []entities.Runner
	for _, value := range result {
		runner := &entities.Runner{
			Name:      value.Name,
			Age:       value.Age,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		runners = append(runners, *runner)
	}

	return runners, err
}

func (s *runnerService) FindById(id int) (*entities.Runner, error) {
	return s.repo.FindById(id)
}

func (s *runnerService) UpdateRunner(id int, name string, age int) (*entities.Runner, error) {
	existingRunner, err := s.repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Runner nor found")
		}
		return nil, err
	}

	if name != "" {
		existingRunner.Name = name
	}

	if age != 0 {
		existingRunner.Age = age
	}

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
