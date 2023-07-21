package repository

import "firebase.google.com/go/auth"

// @ 「FirebaseAuth（認証）」に関する、infraメソッドの集まり（インターフェース）を定義。

type FirebaseAuthRepository interface {
	// 「tokenを検証しtokenの中身を返す」メソッドを定義
	VerifyIDToken(token string) (*auth.Token, error) // auth.Tokenはfirebaseのメソッド。tokenの中身を格納する。
}
