package repository

import "firebase.google.com/go/auth"

type FirebaseAuthRepository interface {
	VerifyIDToken(token string) (*auth.Token, error)
}
