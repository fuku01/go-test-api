package repository

import "github.com/fuku01/go-test-api/app/domain/model"

// !「TodoRepositoryインターフェース（メソッドの集まり）」を定義する。
// ! これらのメソッドは、「infra層」で定義して、「usecase層」で使用（呼び出す）する。

type TodoRepository interface {
	GetAll(userID uint) ([]*model.Todo, error)               // 全てのTodoを取得するメソッドを定義
	Create(content string, userID uint) (*model.Todo, error) // 新しいTodoを作成するメソッドを定義
	Delete(ID uint, userID uint) error                       // 指定したTodoを削除するメソッドを定義
}
