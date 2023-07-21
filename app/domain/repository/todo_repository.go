package repository

import "github.com/fuku01/go-test-api/app/domain/model"

// @ 「Todo」に関する、infraメソッドの集まり（インターフェース）を定義。

type TodoRepository interface {
	// 全てのTodoを取得するメソッドを定義
	GetAll(userID uint) ([]*model.Todo, error)

	// 新しいTodoを作成するメソッドを定義
	Create(content string, userID uint) (*model.Todo, error)

	// 指定したTodoを削除するメソッドを定義
	Delete(ID uint, userID uint) error
}
