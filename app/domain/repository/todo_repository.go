package repository

import "github.com/fuku01/go-test-api/app/domain/model"

type TodoRepository interface {
	GetAll() ([]*model.Todo, error) // GetAllメソッドを定義
}
