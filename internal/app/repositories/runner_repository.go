package repositories

import (
	"RunningTracker/internal/app/entities"

	"gorm.io/gorm"
)

// Define o contrato para o repositório
type RunnerRepository interface {
	Create(runner *entities.Runner) error
}
type runnerRepository struct {
	db *gorm.DB
}

// Cria uma nova instância de runnerRepository com a conexão ao banco
func NewRunnerRepository(db *gorm.DB) RunnerRepository {
	return &runnerRepository{db: db}
}

// Método que usa o GORM para salvar a entidade Runner
func (r *runnerRepository) Create(runner *entities.Runner) error {
	return r.db.Create(runner).Error
}
