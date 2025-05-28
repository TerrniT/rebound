package usecase

import (
	"time"

	"github.com/terrnit/rebound/backend/internal/entity"
)

type authUseCase struct {
	repo     entity.AuthRepository
	userRepo entity.UserRepository
}

// GenerateToken implements entity.AuthUseCase.
func (uc *authUseCase) GenerateToken(user *entity.User) (*entity.Token, error) {
	panic("unimplemented")
}

// RevokeToken implements entity.AuthUseCase.
func (uc *authUseCase) RevokeToken(token string) error {
	panic("unimplemented")
}

// ValidateToken implements entity.AuthUseCase.
func (uc *authUseCase) ValidateToken(token string) (*entity.User, error) {
	panic("unimplemented")
	// tokenEntity, err := uc.repo.GetToken(token)
	// if err != nil {
	// 	return false, err
	// }
	// return tokenEntity != nil && tokenEntity.ExpiresAt.After(time.Now()), nil
}

func NewAuthUseCase(repo entity.AuthRepository, userRepo entity.UserRepository) entity.AuthUseCase {
	return &authUseCase{
		repo:     repo,
		userRepo: userRepo,
	}
}

// func (uc *authUseCase) GenerateToken(user *entity.User) (string, error) {
// 	// TODO: Implement token generation
// 	return "", nil
// }

func (uc *authUseCase) CreateToken(userID int64, token string, expiresAt time.Time) error {
	return uc.repo.CreateToken(userID, token, expiresAt)
}

func (uc *authUseCase) GetToken(token string) (*entity.Token, error) {
	return uc.repo.GetToken(token)
}

func (uc *authUseCase) DeleteToken(token string) error {
	return uc.repo.DeleteToken(token)
}
