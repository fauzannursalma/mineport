package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type NonFormalEducationRepository interface {
	GetAll() ([]entity.FormalEducation, error)
	GetByID(id uint) (entity.FormalEducation, error)
	Create(formalEducation *entity.FormalEducation) error
	Update(formalEducation *entity.FormalEducation) error
	Delete(id uint) error
}

type nonFormalEducationRepository struct {
  db *gorm.DB
}

func NewNonFormalEducationRepository(db *gorm.DB) *nonFormalEducationRepository {
	return &nonFormalEducationRepository{db}
}

func (r *nonFormalEducationRepository) GetAll() ([]entity.FormalEducation, error) {
	var formalEducations []entity.FormalEducation
  return formalEducations, r.db.Find(&formalEducations).Error
}

func (r *nonFormalEducationRepository) GetByID(id uint) (entity.FormalEducation, error) {
	var formalEducation entity.FormalEducation
	return formalEducation, r.db.First(&formalEducation, id).Error
}

func (r *nonFormalEducationRepository) Create(formalEducation *entity.FormalEducation) error {
	return r.db.Create(formalEducation).Error
}

func (r *nonFormalEducationRepository) Update(formalEducation *entity.FormalEducation) error {
	return r.db.Save(formalEducation).Error
}

func (r *nonFormalEducationRepository) Delete(id uint) error {
  return r.db.Delete(&entity.FormalEducation{}, id).Error
}

