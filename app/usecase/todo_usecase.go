package usecase

import (
	"context"

	firebase "firebase.google.com/go/auth"
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// 1. TodoUsecaseインターフェースを定義
// 2. TodoUsecaseインターフェースを実装する構造体は、この3つのメソッドを実装しなければならない

type TodoUsecase interface {
	GetAll(userID uint) ([]*model.Todo, error)               // GetAllメソッドを定義
	Create(content string, userID uint) (*model.Todo, error) // Createメソッドを定義
	Delete(ID uint, userID uint) error                       // Deleteメソッドを定義
	GetUserByToken(ctx context.Context, token string) (*model.User, error)
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
	authClient     *firebase.Client
}

func NewTodoUsecase(todoRepository repository.TodoRepository, authClient *firebase.Client) TodoUsecase {
	return todoUsecase{todoRepository: todoRepository, authClient: authClient}
}

// GetUserByTokenを定義
func (u todoUsecase) GetUserByToken(ctx context.Context, token string) (*model.User, error) {
	firebaseUser, err := u.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	user, err := u.todoRepository.GetUserByFirebaseUID(firebaseUser.UID)
	if err != nil {
		return nil, err
	}
	return user, nil
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
