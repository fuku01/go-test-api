package usecase

import (
	"context"

	firebase "firebase.google.com/go/auth"
	"github.com/fuku01/go-test-api/app/domain/model"
	"github.com/fuku01/go-test-api/app/domain/repository"
)

type UserUsecase interface {
	GetUserByToken(ctx context.Context, token string) (*model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	authClient     *firebase.Client
}

func NewUserUsecase(userRepository repository.UserRepository, authClient *firebase.Client) UserUsecase {
	return &userUsecase{userRepository: userRepository, authClient: authClient}
}

// GetUserByTokenを定義
func (u userUsecase) GetUserByToken(ctx context.Context, token string) (*model.User, error) {
	firebaseUser, err := u.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.GetUserByFirebaseUID(firebaseUser.UID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
