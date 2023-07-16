package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepository{db: db}
}

// GetAllメソッドを定義
func (r TodoRepository) GetAll() ([]*model.Todo, error) {
	var todos []*model.Todo        // このコードではtodosを定義している。todosはmodel.Todoの型を持つ配列。
	err := r.db.Find(&todos).Error // DBから全てのレコードを取得。エラーがあればerrに代入。
	if err != nil {                // エラーがあれば
		return nil, err // エラーがあればエラーを返す
	}
	return todos, nil // エラーがなければtodosを返す
}

// Createメソッドを定義
func (r TodoRepository) Create(content string) (*model.Todo, error) {
	newTodo := &model.Todo{Content: content} // フロントから受け取ったcontentをtodoに代入。
	err := r.db.Create(newTodo).Error        // DBに保存。エラーがあればerrに代入。
	if err != nil {                          // エラーがあれば
		return nil, err // エラーを返す
	}
	return newTodo, nil // エラーがなければnewTodoを返す
}

// Deleteメソッドを定義
func (r TodoRepository) Delete(ID uint) error {
	todo := &model.Todo{}                     // まず、空のTodo構造体を作成
	r.db.First(todo, ID)                      // 次に、データベースから対象のTodoを検索
	err := r.db.Unscoped().Delete(todo).Error // 最後に、該当のTodoを物理的に削除
	if err != nil {
		return err // エラーがあればエラーを返す
	}
	return nil // エラーがなければnilを返す
}
