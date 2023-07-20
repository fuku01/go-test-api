package usecase

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// @ Todoに関する、usecaseメソッドの集まり（インターフェース）を定義。
type TodoUsecase interface {
	GetAll(userID uint) ([]*model.Todo, error)               // 全てのTodoを取得するメソッドを定義
	Create(content string, userID uint) (*model.Todo, error) // 新しいTodoを作成するメソッドを定義
	Delete(ID uint, userID uint) error                       // 指定したTodoを削除するメソッドを定義
}

// @ 構造体の型。
type todoUsecase struct {
	todoRepository repository.TodoRepository
}

// @ /handler層で、この構造体を使用する（呼び出す）ための関数を定義。
func NewTodoUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepository: todoRepository}
}

// @ /repositoryで定義し、/infraで実装した【DBに関する処理】を呼び出し、さらに【具体的な処理】を実装。（今回は、そのまま返すだけ。）

// GetAllメソッド
func (u todoUsecase) GetAll(userID uint) ([]*model.Todo, error) { // GetAllメソッドを定義
	todos, err := u.todoRepository.GetAll(userID) // DBから全てのレコードを取得。エラーがあればerrに代入。
	if err != nil {                               // エラーがあれば
		return nil, err // エラーを返す
	}
	return todos, nil // エラーがなければtodosを返す
}

// Createメソッド
func (u todoUsecase) Create(content string, userID uint) (*model.Todo, error) { // Createメソッドを定義
	todo, err := u.todoRepository.Create(content, userID) // フロントから受け取ったcontentをtodoに代入
	if err != nil {                                       // エラーがあれば
		return nil, err // エラーを返す
	}
	return todo, nil // エラーがなければtodoを返す
}

// Dleteメソッド
func (u todoUsecase) Delete(ID uint, userID uint) error { // Dleteメソッドを定義
	if err := u.todoRepository.Delete(ID, userID); err != nil { // DBから削除。エラーがあればerrに代入。
		return err // エラーを返す
	}
	return nil // エラーがなければnilを返す
}
