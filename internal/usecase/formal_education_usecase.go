package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)

type FormalEducationUsecase interface {
	GetAll() ([]entity.FormalEducation, error)
	GetByID(id uint) (entity.FormalEducation, error)
	Create(formalEducation *entity.FormalEducation) error
	Update(formalEducation *entity.FormalEducation) error
	Delete(id uint) error
}

type formalEducationUsecase struct {
  formalEducationRepository repository.FormalEducationRepository
}

func NewFormalEducationUsecase(formalEducationRepository repository.FormalEducationRepository) *formalEducationUsecase {
	return &formalEducationUsecase{formalEducationRepository}
}

func (u *formalEducationUsecase) GetAll() ([]entity.FormalEducation, error) {
  return u.formalEducationRepository.GetAll()
}

func (u *formalEducationUsecase) GetByID(id uint) (entity.FormalEducation, error) {
  return u.formalEducationRepository.GetByID(id)
}

func (u *formalEducationUsecase) Create(formalEducation *entity.FormalEducation) error {
	return u.formalEducationRepository.Create(formalEducation)
}

func (u *formalEducationUsecase) Update(formalEducation *entity.FormalEducation) error {
	return u.formalEducationRepository.Update(formalEducation)
}

func (u *formalEducationUsecase) Delete(id uint) error {
  return u.formalEducationRepository.Delete(id)
}

