package services

import (
	"RunningTracker/internal/app/entities"
	"RunningTracker/internal/app/repositories"
	"time"
)

// Define o contrato para os serviços que lidam com a lógica do negócio
type RunnerService interface {
	CreateRunner(name string, Age int) (*entities.Runner, error)
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
