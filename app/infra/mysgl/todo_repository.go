package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepositoryメソッドを定義
// 1. DBを引数に取り、TodoRepository構造体を返す
// 2. TodoRepository構造体のフィールドに引数で受け取ったDBを代入
// 3. TodoRepository構造体を返す
// 4. この関数を呼び出すと、DBを引数に取り、TodoRepository構造体を返す関数が作成される

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepository{db: db}
}

// firebaseUIDからuserIDを取得するメソッドを定義
func (r TodoRepository) GetUserByFirebaseUID(firebaseUID string) (*model.User, error) {
	var user model.User
	err := r.db.Where("firebase_uid = ?", firebaseUID).First(&user).Error // DBからfirebase_uidが一致するレコードを取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllメソッドを定義
func (r TodoRepository) GetAll(userID uint) ([]*model.Todo, error) { //user_idを引数に追加
	var todos []*model.Todo
	err := r.db.Where("user_id = ?", userID).Find(&todos).Error // DBからuser_idが一致するレコードを全て取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Createメソッドを定義
func (r TodoRepository) Create(content string, userID uint) (*model.Todo, error) {
	newTodo := &model.Todo{Content: content, UserID: userID} //contentとuser_idを引数にTodo構造体を作成
	err := r.db.Create(newTodo).Error                        // DBに保存。エラーがあればerrに代入。
	if err != nil {                                          // エラーがあれば
		return nil, err // エラーを返す
	}
	return newTodo, nil // エラーがなければnewTodoを返す
}

// Deleteメソッドを定義
func (r TodoRepository) Delete(ID uint, userID uint) error {
	todo := &model.Todo{}                             // まず、空のTodo構造体を作成
	r.db.Where("user_id = ?", userID).First(todo, ID) // 次に、DBからuser_idが一致するレコードを取得
	err := r.db.Unscoped().Delete(todo).Error         // 最後に、取得したレコードを削除
	if err != nil {
		return err // エラーがあればエラーを返す
	}
	return nil // エラーがなければnilを返す
}
