package mysgl

import (
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB // DB接続に必要な情報を格納する。
}

func NewUserRepository(database *gorm.DB) repository.UserRepository {
	return &userRepository{db: database}
}

// GetUserByFirebaseUIDメソッドを定義（firebaseUIDからuserIDを取得するメソッド）
func (r userRepository) GetUserByFirebaseUID(firebaseUID string) (*model.User, error) {
	var user model.User                                                   // User構造体を作成
	err := r.db.Where("firebase_uid = ?", firebaseUID).First(&user).Error // DBから「firebase_uid」が一致するレコードを取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return &user, nil // エラーがなければuserを返す
}
