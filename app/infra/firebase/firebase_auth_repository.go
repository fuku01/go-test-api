package firebase

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

type firebaseAuthRepository struct {
	client *auth.Client
	ctx    context.Context
}

func NewFirebaseAuthRepository(client2 *auth.Client, ctx2 context.Context) repository.FirebaseAuthRepository {
	return &firebaseAuthRepository{client: client2, ctx: ctx2}
}

func (r firebaseAuthRepository) VerifyIDToken(token string) (*auth.Token, error) {
	return r.client.VerifyIDToken(r.ctx, token)
}
