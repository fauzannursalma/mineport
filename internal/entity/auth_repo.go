package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
    CreateUser(user *entity.User) error
    GetUserByEmail(email string) (*entity.User, error)
}

type authRepository struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
    return &authRepository{db}
}

func (r *authRepository) CreateUser(user *entity.User) error {
    return r.db.Create(user).Error
}

func (r *authRepository) GetUserByEmail(email string) (*entity.User, error) {
    var user entity.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
