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

// GetAllWithTagsメソッド
func (r todoRepository) GetAllWithTags(userID uint) ([]*model.Todo, error) {
	var todos []*model.Todo // Todo構造体の配列を作成

	err := r.db.Preload("Tags").Where("user_id = ?", userID).Find(&todos).Error // Preloadとは、指定した関連付けを事前に読み込むことができるメソッド。ここでは、Todoに紐づくTagを事前に読み込んでいる。
	if err != nil {
		return nil, err
	}

	return todos, nil // エラーがなければtodosを返す
}

// ! CreateWithTagsメソッド（トランザクションを使用して、TodoとTagを同時に作成）
func (r todoRepository) CreateWithTags(content string, userID uint, tagNames []string) (*model.Todo, error) { // contentとuserIDとtagNamesを引数に追加し、戻り値を*model.Todoにする。
	newTodo := &model.Todo{Content: content, UserID: userID} // contentとuser_idを引数にTodo構造体を作成
	tx := r.db.Begin()                                       // 「Begin()」とは、トランザクションを開始するメソッド。（開始すると、以降の処理は全てトランザクション内で実行されるため、エラーがあればロールバック（処理を元に戻す）される。）

	// 1. Todoを作成
	err := tx.Create(newTodo).Error //「tx.Create」とは、トランザクション内でDBに保存するメソッド。
	if err != nil {
		tx.Rollback() // エラーがあればロールバック（処理を元に戻す）
		return nil, err
	}
	// 2. Tagを作成
	newTags := []model.Tag{}        // Tag構造体の配列を作成
	for _, name := range tagNames { // tagNames（Tag構造体の配列）をループさせる
		newTags = append(newTags, model.Tag{Name: name}) // appendとは、[配列]に要素を追加するメソッド。（追加先の[配列]と追加する要素。）
	}
	// 3. TodoとTagを関連付ける
	err = tx.Model(newTodo).Association("Tags").Append(newTags) // 「tx.Model(newTodo).Association("Tags")」とは、TodoとTagを関連付けるメソッド。「Append(newTags)」とは、関連付けるTagを追加するメソッド。
	if err != nil {
		tx.Rollback() // エラーがあればロールバック
		return nil, err
	}

	err = tx.Commit().Error // 「tx.Commit()」とは、トランザクションをコミット(処理を確定)するメソッド。
	if err != nil {
		tx.Rollback() // エラーがあればロールバック（処理を元に戻す）
		return nil, err
	}

	return newTodo, nil // エラーがなければnewTodo（新しく作成したTodo）を返す

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

// ! DeleteWithTagsメソッド（トランザクションを使用して、Todoとそれに紐付くTagを全て削除する）
func (r todoRepository) DeleteWithTags(ID uint, userID uint) error {
	todo := &model.Todo{} // 空のTodo構造体を作成
	tx := r.db.Begin()    // トランザクションを開始

	// 1.Todoを取得
	err := r.db.Where("user_id = ?", userID).First(todo, ID).Error
	if err != nil {
		tx.Rollback() // エラーがあればロールバック
		return err
	}
	// 2.Todoに紐づくTagを全て削除する
	err = tx.Where("todo_id = ?", todo.ID).Unscoped().Delete(&model.Tag{}).Error
	if err != nil {
		tx.Rollback() // エラーがあればロールバック
		return err
	}
	// 3.Todoを削除
	err = tx.Unscoped().Delete(todo).Error
	if err != nil {
		tx.Rollback() // エラーがあればロールバック
		return err
	}

	err = tx.Commit().Error // トランザクションをコミット
	if err != nil {
		tx.Rollback() // エラーがあればロールバック
		return err
	}

	return nil // エラーがなければnilを返す
}

// Deleteメソッド
func (r todoRepository) Delete(ID uint, userID uint) error {
	todo := &model.Todo{}                             // まず、空のTodo構造体を作成
	r.db.Where("user_id = ?", userID).First(todo, ID) // 次に、DBからuser_idが一致するレコードを取得
	err := r.db.Unscoped().Delete(todo).Error         // 最後に、取得したレコードを削除。(Unscoped()とは、物理削除を行うメソッド。)
	if err != nil {
		return err
	}

	return nil // エラーがなければnilを返す
}

// ! EditWithTagsメソッド(トランザクションを使用して、Todoとそれに紐付くTagの追加と削除を行う)
func (r todoRepository) EditWithTags(ID uint, userID uint, content string, addTagNames []string, deleteTagIDs []uint) (*model.Todo, error) {
	// トランザクションを開始します。
	tx := r.db.Begin()

	// 現在のTodoを取得します。
	todo := &model.Todo{}
	err := tx.Where("user_id = ?", userID).First(todo, ID).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// contentが空でなければ、contentを更新します。
	if content != "" {
		todo.Content = content
		err = tx.Save(todo).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 追加すべき新しいタグがあれば、それをTodoに追加します。
	if len(addTagNames) > 0 {
		for _, name := range addTagNames { // addTagNamesの要素を1つずつ取り出す
			err = tx.Create(&model.Tag{Name: name, Todo: todo}).Error // Todoに紐付くTagを作成。
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// 削除すべきタグがあれば、それをTodoから削除します。
	if len(deleteTagIDs) > 0 {
		for _, id := range deleteTagIDs { // deleteTagIDsの要素を1つずつ取り出す

			tag := &model.Tag{}                                          // 空のTag構造体を作成
			err := tx.Where("todo_id = ?", todo.ID).First(tag, id).Error // todo_idが一致するレコードを取得
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			err = tx.Unscoped().Delete(tag).Error // 取得したレコードを削除
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// すべての操作が正常に終了したら、トランザクションをコミットします。
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 最終的に更新されたTodoを返します。
	return todo, nil
}
