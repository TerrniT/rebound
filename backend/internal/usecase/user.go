package usecase

import (
	"github.com/terrnit/rebound/backend/internal/entity"
)

type userUseCase struct {
	repo entity.UserRepository
}

// Login implements entity.UserUseCase.
func (uc *userUseCase) Login(email string, password string) (string, error) {
	panic("unimplemented")
}

// Register implements entity.UserUseCase.
func (uc *userUseCase) Register(email string, password string, name string) (*entity.User, error) {
	panic("unimplemented")
}

// ValidateToken implements entity.UserUseCase.
func (uc *userUseCase) ValidateToken(token string) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserUseCase(repo entity.UserRepository) entity.UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(user *entity.User) error {
	return uc.repo.Create(user)
}

func (uc *userUseCase) GetUser(id int64) (*entity.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *userUseCase) GetUserByEmail(email string) (*entity.User, error) {
	return uc.repo.GetByEmail(email)
}

func (uc *userUseCase) UpdateUser(user *entity.User) error {
	return uc.repo.Update(user)
}

func (uc *userUseCase) DeleteUser(id int64) error {
	return uc.repo.Delete(id)
}
