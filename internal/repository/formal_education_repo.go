package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type FormalEducationRepository interface {
	GetAll() ([]entity.FormalEducation, error)
	GetByID(id uint) (entity.FormalEducation, error)
	Create(formalEducation *entity.FormalEducation) error
	Update(formalEducation *entity.FormalEducation) error
	Delete(id uint) error
}

type formalEducationRepository struct {
  db *gorm.DB
}

func NewFormalEducationRepository(db *gorm.DB) *formalEducationRepository {
  return &formalEducationRepository{db}
}

func (r *formalEducationRepository) GetAll() ([]entity.FormalEducation, error) {
	var formalEducations []entity.FormalEducation
	return formalEducations, r.db.Find(&formalEducations).Error
}

func (r *formalEducationRepository) GetByID(id uint) (entity.FormalEducation, error) {
	var formalEducation entity.FormalEducation
	return formalEducation, r.db.First(&formalEducation, id).Error
}

func (r *formalEducationRepository) Create(formalEducation *entity.FormalEducation) error {
  return r.db.Create(formalEducation).Error
}

func (r *formalEducationRepository) Update(formalEducation *entity.FormalEducation) error {
  return r.db.Save(formalEducation).Error
}

func (r *formalEducationRepository) Delete(id uint) error {
  return r.db.Delete(&entity.FormalEducation{}, id).Error
}

