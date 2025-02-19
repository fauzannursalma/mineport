package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type WorkExperienceRepository interface {
	GetAll() ([]entity.WorkExperience, error)
	GetByID(id uint) (entity.WorkExperience, error)
	Create(workExperience *entity.WorkExperience) error
	Update(workExperience *entity.WorkExperience) error
	Delete(id uint) error
}

type workExperienceRepository struct {
  db *gorm.DB
}

func NewWorkExperienceRepository(db *gorm.DB) *workExperienceRepository {
	return &workExperienceRepository{db}
}

func (r *workExperienceRepository) GetAll() ([]entity.WorkExperience, error) {
	var workExperiences []entity.WorkExperience
	return workExperiences, r.db.Find(&workExperiences).Error
}

func (r *workExperienceRepository) GetByID(id uint) (entity.WorkExperience, error) {
	var workExperience entity.WorkExperience
	return workExperience, r.db.First(&workExperience, id).Error
}

func (r *workExperienceRepository) Create(workExperience *entity.WorkExperience) error {
	return r.db.Create(workExperience).Error
}

func (r *workExperienceRepository) Update(workExperience *entity.WorkExperience) error {
  return r.db.Save(workExperience).Error
}

func (r *workExperienceRepository) Delete(id uint) error {
	return r.db.Delete(&entity.WorkExperience{}, id).Error
}
