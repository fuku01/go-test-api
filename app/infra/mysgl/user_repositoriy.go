package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

// @ 構造体の型。
type userRepository struct {
	db *gorm.DB
}

// @ /usecase層で、この構造体を使用する（呼び出す）ための関数を定義。
// ?　1.引数：db *gorm.DB　= DB接続に必要な情報を格納する。
// ?　2.戻り値の型：repository.UserRepository　= 「repository/user_repository.go」で定義したUserRepositoryインターフェース。
// ?　3.戻り値：&UserRepository{db: db}　= UserRepository構造体（type UserRepositoryで定義）を返す。
func NewUserRepository(database *gorm.DB) repository.UserRepository {
	return &userRepository{db: database}
}

// @ /repositoryで定義したメソッドの、DBに関する処理を実装。
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
