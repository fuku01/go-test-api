package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB // DB接続に必要な情報を格納する。
}

// ! 「usecase層」でこのファイルのメソッドを使用するために、「NewTodoRepositoryメソッド」を定義する。
// ?　1.引数：db *gorm.DB　= DB接続に必要な情報を格納する。
// ?　2.戻り値の型：repository.TodoRepository　= 「repository/todo_repository.go」で定義したTodoRepositoryインターフェース。
// ?　3.戻り値：&TodoRepository{db: db}　= TodoRepository構造体（type TodoRepositoryで定義）のポインタを返す。
func NewTodoRepository(database *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: database}
}

// ! 「domain層」の「repository/todo_repository.go」で定義した「TodoRepositoryインターフェース」のメソッドを実装する。

// GetAllメソッドを定義
func (r todoRepository) GetAll(userID uint) ([]*model.Todo, error) { //user_idを引数に追加。
	var todos []*model.Todo                                     // .Todo構造体の配列を作成
	err := r.db.Where("user_id = ?", userID).Find(&todos).Error // DBからuser_idが一致するレコードを全て取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return todos, nil // エラーがなければtodosを返す
}

// Createメソッドを定義
func (r todoRepository) Create(content string, userID uint) (*model.Todo, error) {
	newTodo := &model.Todo{Content: content, UserID: userID} // contentとuser_idを引数にTodo構造体を作成
	err := r.db.Create(newTodo).Error                        // DBに保存。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return newTodo, nil // エラーがなければnewTodoを返す
}

// Deleteメソッドを定義
func (r todoRepository) Delete(ID uint, userID uint) error {
	todo := &model.Todo{}                             // まず、空のTodo構造体を作成
	r.db.Where("user_id = ?", userID).First(todo, ID) // 次に、DBからuser_idが一致するレコードを取得
	err := r.db.Unscoped().Delete(todo).Error         // 最後に、取得したレコードを削除。エラーがあればerrに代入。
	if err != nil {
		return err
	}
	return nil // エラーがなければnilを返す
}
