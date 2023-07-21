package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

// @ 構造体の型。
type userRepository struct {
	db *gorm.DB // 型：gorm.DB（db接続に必要な情報を格納する際に使用する型）
}

// @ /usecase層で、この構造体を使用する（呼び出す）ための関数を定義。
// ? db2に引数でデータベース接続情報を受け取り、repository.UserRepositoryのインターフェースを満たすような新しいuserRepository構造体を作成して返す。
func NewUserRepository(db2 *gorm.DB) repository.UserRepository {
	return &userRepository{db: db2}
}

// @ /repositoryで定義したメソッドの【DBに関する処理】を実装。

// GetUserByFirebaseUIDメソッド　（※「firebaseUID」から「user」を取得するメソッド）
func (r userRepository) GetUserByFirebaseUID(firebaseUID string) (*model.User, error) {
	var user model.User                                                   // User構造体を作成
	err := r.db.Where("firebase_uid = ?", firebaseUID).First(&user).Error // DBから「firebase_uid」が一致する「user」を取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return &user, nil // エラーがなければuserを返す
}

// GetLoginUserメソッド
func (r userRepository) GetLoginUser(userID uint) (*model.User, error) {
	var user model.User                                    // User構造体を作成
	err := r.db.Where("id = ?", userID).First(&user).Error // DBから「id」が一致する「user」を取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return &user, nil // エラーがなければuserを返す
}
