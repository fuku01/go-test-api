package firebase

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

// @ 構造体の型。
type firebaseAuthRepository struct {
	client *auth.Client
}

// @ /usecase層で、この構造体を使用する（呼び出す）ための関数を定義。
// ? client2,ctx2に引数で各々のインターフェースを満たすオブジェクトを受け取り、FirebaseAuthRepositoryのインターフェースを満たすような新しいfirebaseAuthRepository構造体を作成して返す。
func NewFirebaseAuthRepository(client2 *auth.Client) repository.FirebaseAuthRepository {
	return &firebaseAuthRepository{client: client2}
}

// @ /repositoryで定義した【認証に関する処理】を実装。
// （「tokenを検証しtokenの中身を返す」メソッド）
func (r firebaseAuthRepository) VerifyIDToken(token string) (*auth.Token, error) { // ?auth.Tokenはfirebaseのメソッド。tokenの中身を格納する。
	ctx := context.Background()
	return r.client.VerifyIDToken(ctx, token) // ?「VerifyIDToken」はfirebaseのメソッド。tokenを検証しtokenの中身を返す。
}
