package usecase

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// @ 「Todo」に関する、usecaseメソッドの集まり（インターフェース）を定義。
type TodoUsecase interface {
	// 全てのTodoを取得するメソッドを定義
	GetAll(token string) ([]*model.Todo, error)

	// 新しいTodoを作成するメソッドを定義
	Create(content string, token string) (*model.Todo, error)

	// 指定したTodoを削除するメソッドを定義
	Delete(ID uint, token string) error

	// TagテーブルとTodoテーブルを結合して、全てのTodoに加えて、そのTodoが持つ全てのTagを取得するメソッドを定義
	GetAllWithTags(token string) ([]*model.Todo, error)

	// *CreateWithTagsメソッド（トランザクションを使用して、TodoとTagを同時に作成）
	CreateWithTags(content string, token string, tagNames []string) (*model.Todo, error)
}

// @ 構造体の型。
type todoUsecase struct {
	tr  repository.TodoRepository
	ur  repository.UserRepository
	far repository.FirebaseAuthRepository
}

// @ /handler層で、この構造体を使用する（呼び出す）ための関数を定義。
// ? tr2,ur2,far2に引数で各々のインターフェースを満たすオブジェクトを受け取り、TodoUsecaseのインターフェースを満たすような新しいtodoUsecase構造体を作成して返す。
func NewTodoUsecase(tr2 repository.TodoRepository, ur2 repository.UserRepository, far2 repository.FirebaseAuthRepository) TodoUsecase {
	return &todoUsecase{tr: tr2, ur: ur2, far: far2}
}

// @ /repositoryで定義し、/infraで実装した【DBに関する処理】を呼び出し、さらに【具体的な処理】を実装。（今回は、そのまま返すだけ。）

// GetAllメソッド
func (u todoUsecase) GetAll(token string) ([]*model.Todo, error) {

	firebaseUser, err := u.far.VerifyIDToken(token) // トークンを検証
	if err != nil {
		return nil, err
	}
	user, err := u.ur.GetUserByFirebaseUID(firebaseUser.UID) // ユーザーを取得
	if err != nil {
		return nil, err
	}
	todos, err := u.tr.GetAll(user.ID) // DBから全てのレコードを取得。エラーがあればerrに代入。
	if err != nil {                    // エラーがあれば
		return nil, err // エラーを返す
	}
	return todos, nil // エラーがなければtodosを返す
}

// GetAllWithTagsメソッド
func (u todoUsecase) GetAllWithTags(token string) ([]*model.Todo, error) {

	firebaseUser, err := u.far.VerifyIDToken(token) // トークンを検証
	if err != nil {
		return nil, err
	}
	user, err := u.ur.GetUserByFirebaseUID(firebaseUser.UID) // ユーザーを取得
	if err != nil {
		return nil, err
	}
	todosWithTags, err := u.tr.GetAllWithTags(user.ID) // DBから全てのレコードを取得。エラーがあればerrに代入。
	if err != nil {                                    // エラーがあれば
		return nil, err // エラーを返す
	}
	return todosWithTags, nil // エラーがなければtodosWithTagsを返す
}

// *CreateWithTagsメソッド
func (u todoUsecase) CreateWithTags(content string, token string, tagNames []string) (*model.Todo, error) {

	firebaseUser, err := u.far.VerifyIDToken(token) // トークンを検証
	if err != nil {
		return nil, err
	}
	user, err := u.ur.GetUserByFirebaseUID(firebaseUser.UID) // ユーザーを取得
	if err != nil {
		return nil, err
	}
	newTodo, err := u.tr.CreateWithTags(content, user.ID, tagNames) // DBに保存。エラーがあればerrに代入。
	if err != nil {                                                 // エラーがあれば
		return nil, err // エラーを返す
	}
	return newTodo, nil // エラーがなければtodoを返す
}

// Createメソッド
func (u todoUsecase) Create(content string, token string) (*model.Todo, error) {

	firebaseUser, err := u.far.VerifyIDToken(token) // トークンを検証
	if err != nil {
		return nil, err
	}
	user, err := u.ur.GetUserByFirebaseUID(firebaseUser.UID) // ユーザーを取得
	if err != nil {
		return nil, err
	}
	newTodo, err := u.tr.Create(content, user.ID) // フロントから受け取ったcontentをtodoに代入。
	if err != nil {                               // エラーがあれば
		return nil, err // エラーを返す
	}
	return newTodo, nil // エラーがなければtodoを返す
}

// Dleteメソッド
func (u todoUsecase) Delete(ID uint, token string) error {

	firebaseUser, err := u.far.VerifyIDToken(token) // トークンを検証
	if err != nil {
		return err
	}
	user, err := u.ur.GetUserByFirebaseUID(firebaseUser.UID) // ユーザーを取得
	if err != nil {
		return err
	}

	if err := u.tr.Delete(ID, user.ID); err != nil { // DBから削除。エラーがあればerrに代入。
		return err // エラーを返す
	}
	return nil // エラーがなければnilを返す
}
