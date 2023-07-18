package repository

import "github.com/fuku01/go-test-api/app/domain/model"

type TodoRepository interface {
	GetAll() ([]*model.Todo, error)             // GetAllメソッドを定義
	Create(content string) (*model.Todo, error) // Createメソッドを定義
	Delete(ID uint) error                       // Deleteメソッドを定義
}
