package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
