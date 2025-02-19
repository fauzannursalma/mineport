package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)

type UserUsecase interface {
	GetAll() ([]entity.User, error)
	GetByID(id uint) (entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
  return &userUsecase{userRepo}
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.userRepository.GetAll()
}

func (u *userUsecase) GetByID(id uint) (entity.User, error) {
  return u.userRepository.GetByID(id)
}

func (u *userUsecase) Create(user *entity.User) error {
	return u.userRepository.Create(user)
}

func (u *userUsecase) Update(user *entity.User) error {
	return u.userRepository.Update(user)
}

func (u *userUsecase) Delete(id uint) error {
	return u.userRepository.Delete(id)
}