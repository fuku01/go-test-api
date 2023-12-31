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

	// TagテーブルとTodoテーブルを結合して、全てのTodoに加えて、そのTodoが持つ全てのTagを取得するメソッドを定義
	GetAllWithTags(userID uint) ([]*model.Todo, error)

	// *CreateWithTagsメソッド（トランザクションを使用して、TodoとTagを同時に作成）
	CreateWithTags(content string, userID uint, tagNames []string) (*model.Todo, error)

	// *DeleteWithTagsメソッド（トランザクションを使用して、TodoとTagを同時に削除）
	DeleteWithTags(ID uint, userID uint) error

	// !EditWithTagsメソッド(トランザクションを使用して、Todoとそれに紐付くTagの追加と削除を行う)
	EditWithTags(ID uint, userID uint, content string, addTagNames []string, deleteTagIDs []uint) (*model.Todo, error)
}
