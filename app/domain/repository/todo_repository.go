package repository

import "github.com/fuku01/go-test-api/app/domain/model"

// 1. TodoRepositoryインターフェースを定義
// 2. TodoRepositoryインターフェースを実装する構造体は、この3つのメソッドを実装しなければならない

type TodoRepository interface {
	GetAll(userID uint) ([]*model.Todo, error)               //GetAllメソッドを定義
	Create(content string, userID uint) (*model.Todo, error) //Createメソッドを定義
	Delete(ID uint, userID uint) error                       //Deleteメソッドを定義
	GetUserByFirebaseUID(firebaseUID string) (*model.User, error)
}
