package usecase

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// ! 1.このファイルで使用する「TodoUsecaseインターフェース（メソッドの集まり）」を定義する。
type TodoUsecase interface {
	GetAll(userID uint) ([]*model.Todo, error)
	Create(content string, userID uint) (*model.Todo, error)
	Delete(ID uint, userID uint) error
}

// ! 2.「handler層」の「todo_handler.go」で使用する「TodoUsecaseストラクト（構造体）」を定義する。
type todoUsecase struct {
	todoRepository repository.TodoRepository
}

// ! 3.「handler層」でこのファイルのメソッドを使用するために、「NewTodoUsecaseメソッド」を定義する。
func NewTodoUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepository: todoRepository}
}

// GetAllメソッドを定義
func (u todoUsecase) GetAll(userID uint) ([]*model.Todo, error) { // GetAllメソッドを定義
	todos, err := u.todoRepository.GetAll(userID) // DBから全てのレコードを取得。エラーがあればerrに代入。
	if err != nil {                               // エラーがあれば
		return nil, err // エラーを返す
	}
	return todos, nil // エラーがなければtodosを返す
}

// Createメソッドを定義
func (u todoUsecase) Create(content string, userID uint) (*model.Todo, error) { // Createメソッドを定義
	todo, err := u.todoRepository.Create(content, userID) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                                       // エラーがあれば
		return nil, err // エラーを返す
	}
	return todo, nil // エラーがなければtodoを返す
}

// Dleteメソッドを定義
func (u todoUsecase) Delete(ID uint, userID uint) error { // Dleteメソッドを定義
	if err := u.todoRepository.Delete(ID, userID); err != nil { // DBから削除。エラーがあればerrに代入。
		return err // エラーを返す
	}
	return nil // エラーがなければnilを返す
}
