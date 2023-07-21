package usecase

import (
	"context"

	firebase "firebase.google.com/go/auth"
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// @ 「User」に関する、usecaseメソッドの集まり（インターフェース）を定義。
type UserUsecase interface {
	// 「token」から「firebaseUser」を取得し、さらに「firebaseUser.UID」から「user」を取得するメソッドを定義
	GetUserByToken(ctx context.Context, token string) (*model.User, error)

	// ログイン中のユーザー情報をDBから取得するメソッドを定義
	GetLoginUser(userID uint) (*model.User, error)
}

// @ 構造体の型。
type userUsecase struct {
	ur         repository.UserRepository
	authClient *firebase.Client // 「firebase.Client」はVerifyIDTokenメソッドの結果を格納するための型。
}

// @ /handler層で、この構造体を使用する（呼び出す）ための関数を定義。
// ? ur2,authClient2に引数で各々のインターフェースを満たすオブジェクトを受け取り、UserUsecaseのインターフェースを満たすような新しいuserUsecase構造体を作成して返す。
func NewUserUsecase(ur2 repository.UserRepository, authClient2 *firebase.Client) UserUsecase {
	return &userUsecase{ur: ur2, authClient: authClient2}
}

// @ /repositoryで定義し、/infraで実装した【DBに関する処理】を使用し（呼び出し）、さらに【具体的な処理】を実装。

// GetUserByTokenメソッド
func (u userUsecase) GetUserByToken(ctx context.Context, token string) (*model.User, error) {
	// *「token」から「firebaseUser」を取得。
	firebaseUser, err := u.authClient.VerifyIDToken(ctx, token) // ?「VerifyIDToken」はfirebaseのメソッド。tokenを検証しtokenの中身を返す。）
	if err != nil {
		return nil, err
	}
	// *「firebaseUser.UID」から「user」を取得。
	user, err := u.ur.GetUserByFirebaseUID(firebaseUser.UID) // ?「.UID」は、firebaseのメソッド。firebaseUserのUIDを返す。）
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetLoginUserメソッド
func (u userUsecase) GetLoginUser(userID uint) (*model.User, error) {
	user, err := u.ur.GetLoginUser(userID) // DBから「id」が一致する「user」を取得。エラーがあればerrに代入。
	if err != nil {
		return nil, err
	}
	return user, nil // エラーがなければuserを返す
}
