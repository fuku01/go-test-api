package repository

import "github.com/fuku01/go-test-api/app/domain/model"

// @ 「User」に関する、メソッドの集まり（インターフェース）を定義。

type UserRepository interface {
	// 「firebaseUID」から「user」を取得するメソッドを定義
	GetUserByFirebaseUID(firebaseUID string) (*model.User, error)

	// ログイン中のユーザー情報をDBから取得するメソッドを定義
	GetLoginUser(userID uint) (*model.User, error)
}
