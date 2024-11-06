package repositories

import (
	"RunningTracker/internal/app/entities"

	"gorm.io/gorm"
)

// Define o contrato para o repositório
type RunnerRepository interface {
	Create(runner *entities.Runner) error
	Update(runner *entities.Runner) error
	Delete(id uint) error
}
type runnerRepository struct {
	db *gorm.DB
}

// Cria uma nova instância de runnerRepository com a conexão ao banco
func NewRunnerRepository(db *gorm.DB) RunnerRepository {
	return &runnerRepository{db: db}
}

// Método que usa o GORM para salvar a entidade Runner
func (repo *runnerRepository) Create(runner *entities.Runner) error {
	return repo.db.Create(runner).Error
}

func (repo *runnerRepository) Update(runner *entities.Runner) error {
	return repo.db.Model(&runner).Updates(runner).Error
}

func (repo *runnerRepository) Delete(id uint) error {
	return repo.db.Delete(&entities.Runner{}, id).Error
}
