package usecase

import (
	"context"

	firebase "firebase.google.com/go/auth"
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// @ Userに関する、usecaseメソッドの集まり（インターフェース）を定義。
type UserUsecase interface {
	GetUserByToken(ctx context.Context, token string) (*model.User, error) // *「token」から「firebaseUser」を取得し、さらに「firebaseUser.UID」から「user」を取得するメソッドを定義
}

// @ 構造体の型。
type userUsecase struct {
	userRepository repository.UserRepository
	authClient     *firebase.Client // ?「firebase.Client」は firebaseのメソッドを使用するために必要な情報を格納する。
}

// @ /handler層で、この構造体を使用する（呼び出す）ための関数を定義。
func NewUserUsecase(userRepository repository.UserRepository, authClient *firebase.Client) UserUsecase {
	return &userUsecase{userRepository: userRepository, authClient: authClient}
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
	user, err := u.userRepository.GetUserByFirebaseUID(firebaseUser.UID) // ?「.UID」は、firebaseのメソッド。firebaseUserのUIDを返す。）
	if err != nil {
		return nil, err
	}
	return user, nil
}
