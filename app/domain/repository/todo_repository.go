package repository

import "github.com/fuku01/go-test-api/app/domain/model"

// @ 「Todo」に関する、メソッドの集まり（インターフェース）を定義。
// ? 1.「infra層」で【DBに関する処理】を実装。
// ? 2.「usecase層」で、その処理を使用（呼び出す）して、さらに【具体的な処理】を実装。
// ? 3.「handler層」で、フロントからのHTTPリクエストを受け取り、対応するusecase層の処理を呼び出し、フロントに返すレスポンスを生成する。
// ? 4.「/main.go」で、handler層の処理をルーティング（URLと紐付け）する。

type TodoRepository interface {
	GetAll(userID uint) ([]*model.Todo, error)               // 全てのTodoを取得するメソッドを定義
	Create(content string, userID uint) (*model.Todo, error) // 新しいTodoを作成するメソッドを定義
	Delete(ID uint, userID uint) error                       // 指定したTodoを削除するメソッドを定義
}
