package usecase

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

type TodoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return TodoUsecase{todoRepository: todoRepository}
}
func (u TodoUsecase) GetAll() ([]*model.Todo, error) {
	todos, err := u.todoRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
