package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(id uint) (entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("FormalEducations").
	Preload("NonFormalEducations").
	Preload("WorkExperiences").
	Preload("Skills").
	Preload("Projects").
	Preload("Certificates").
	Find(&users).Error

	return users, err
}

func (r *userRepository) GetByID(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.Preload("FormalEducations").
	Preload("NonFormalEducations").
	Preload("WorkExperiences").
	Preload("Skills").
	Preload("Projects").
	Preload("Certificates").
	Where("id = ?", id).
	First(&user).Error

	return user, err
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}