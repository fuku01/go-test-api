package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

// @ 構造体の型。
type todoRepository struct {
	db *gorm.DB // 型：gorm.DB（db接続に必要な情報を格納する際に使用する型）
}

// @ /usecase層で、この構造体を使用する（呼び出す）ための関数を定義。
// ? db2に引数でデータベース接続情報を受け取り、repository.TodoRepositoryのインターフェースを満たすような新しいtodoRepository構造体を作成して返す。
func NewTodoRepository(db2 *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db2} // 構造体を返す。(db2とは、引数で受け取ったdb *gorm.DBのこと。）
}

// @ /repositoryで定義したメソッドの【DBに関する処理】を実装。

// GetAllメソッド
func (r todoRepository) GetAll(userID uint) ([]*model.Todo, error) { //user_idを引数に追加。
	var todos []*model.Todo // Todo構造体の配列を作成

	err := r.db.Where("user_id = ?", userID).Find(&todos).Error // DBからuser_idが一致するレコードを全て取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}

	return todos, nil // エラーがなければtodosを返す
}

// Createメソッド
func (r todoRepository) Create(content string, userID uint) (*model.Todo, error) {
	newTodo := &model.Todo{Content: content, UserID: userID} // contentとuser_idを引数にTodo構造体を作成
	err := r.db.Create(newTodo).Error                        // DBに保存。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}

	return newTodo, nil // エラーがなければnewTodo（新しく作成したTodo）を返す
}

// Deleteメソッド
func (r todoRepository) Delete(ID uint, userID uint) error {
	todo := &model.Todo{}                             // まず、空のTodo構造体を作成
	r.db.Where("user_id = ?", userID).First(todo, ID) // 次に、DBからuser_idが一致するレコードを取得
	err := r.db.Unscoped().Delete(todo).Error         // 最後に、取得したレコードを削除。エラーがあればerrに代入。
	if err != nil {
		return err
	}

	return nil // エラーがなければnilを返す
}
