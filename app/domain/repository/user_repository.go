package repository

import "github.com/fuku01/go-test-api/app/domain/model"

type UserRepository interface {
	GetUserByFirebaseUID(firebaseUID string) (*model.User, error) // firebaseUIDからuserIDを取得するメソッドを定義
}
