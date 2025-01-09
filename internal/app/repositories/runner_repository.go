package repositories

import (
	"RunningTracker/internal/app/entities"

	"gorm.io/gorm"
)

// Define o contrato para o repositório
type RunnerRepository interface {
	Create(runner *entities.Runner) error
	GetUsers() ([]entities.Runner, error)
	Update(runner *entities.Runner) error
	Delete(id int) error
	FindById(id int) (*entities.Runner, error)
	GetAllRunsByUser(id int) ([]entities.Runner, error)
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

func (repo *runnerRepository) GetUsers() ([]entities.Runner, error) {
	var runner []entities.Runner
	err := repo.db.Find(&runner).Error
	if err != nil {
		return nil, err
	}
	return runner, nil
}

func (repo *runnerRepository) Update(runner *entities.Runner) error {
	err := repo.db.Model(&entities.Runner{}).Updates(runner).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *runnerRepository) Delete(id int) error {
	return repo.db.Delete(&entities.Runner{}, id).Error
}

func (repo *runnerRepository) FindById(id int) (*entities.Runner, error) {
	var runner entities.Runner
	err := repo.db.First(&runner, id).Error
	if err != nil {
		return nil, err
	}
	return &runner, nil
}

func (repo *runnerRepository) GetAllRunsByUser(id int) ([]entities.Runner, error) {
	var runners []entities.Runner
	err := repo.db.Model(&entities.Runner{}).Preload("Races").Find(&runners, id).Error
	return runners, err
}
