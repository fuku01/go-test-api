package usecase

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

type TodoUsecase interface {
	GetAll() ([]*model.Todo, error)             // GetAllメソッドを定義
	Create(content string) (*model.Todo, error) // Createメソッドを定義
	Delete(ID uint) error                       // Deleteメソッドを定義
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return todoUsecase{todoRepository: todoRepository}
}

// GetAllメソッドを定義
func (u todoUsecase) GetAll() ([]*model.Todo, error) { // GetAllメソッドを定義
	todos, err := u.todoRepository.GetAll() // DBから全てのレコードを取得。エラーがあればerrに代入。
	if err != nil {                         // エラーがあれば
		return nil, err // エラーを返す
	}
	return todos, nil // エラーがなければtodosを返す
}

// Createメソッドを定義
func (u todoUsecase) Create(content string) (*model.Todo, error) { // Createメソッドを定義
	todo, err := u.todoRepository.Create(content) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                               // エラーがあれば
		return nil, err // エラーを返す
	}
	return todo, nil // エラーがなければtodoを返す
}

// Dleteメソッドを定義
func (u todoUsecase) Delete(ID uint) error { // Dleteメソッドを定義
	if err := u.todoRepository.Delete(ID); err != nil { // DBから削除。エラーがあればerrに代入。
		return err // エラーを返す
	}
	return nil // エラーがなければnilを返す
}
