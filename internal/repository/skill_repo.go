package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type SkillRepository interface {
	GetAll() ([]entity.Skill, error)
	GetByID(id uint) (entity.Skill, error)
	Create(skill *entity.Skill) error
	Update(skill *entity.Skill) error
	Delete(id uint) error
}

type skillRepository struct {
  db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) *skillRepository {
	return &skillRepository{db}
}

func (r *skillRepository) GetAll() ([]entity.Skill, error) {
	var skills []entity.Skill
	err := r.db.Find(&skills).Error

	return skills, err
}

func (r *skillRepository) GetByID(id uint) (entity.Skill, error) {
	var skill entity.Skill
	err := r.db.Where("id = ?", id).First(&skill).Error

	return skill, err
}

func (r *skillRepository) Create(skill *entity.Skill) error {
	return r.db.Create(skill).Error
}

func (r *skillRepository) Update(skill *entity.Skill) error {
  return r.db.Save(skill).Error
}

func (r *skillRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Skill{}, id).Error
}