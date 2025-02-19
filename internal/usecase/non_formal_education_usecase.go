package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)

type NonFormalEducationUsecase interface {
	GetAll() ([]entity.NonFormalEducation, error)
	GetByID(id uint) (entity.NonFormalEducation, error)
	Create(nonFormalEducation *entity.NonFormalEducation) error
	Update(nonFormalEducation *entity.NonFormalEducation) error
	Delete(id uint) error
}

type nonFormalEducationUsecase struct {
  formalEducationRepository repository.FormalEducationRepository
}

func NewNonFormalEducationUsecase(formalEducationRepository repository.FormalEducationRepository) *nonFormalEducationUsecase {
	return &nonFormalEducationUsecase{formalEducationRepository}
}

func (u *nonFormalEducationUsecase) GetAll() ([]entity.FormalEducation, error) {
  return u.formalEducationRepository.GetAll()
}

func (u *nonFormalEducationUsecase) GetByID(id uint) (entity.FormalEducation, error) {
  return u.formalEducationRepository.GetByID(id)
}

func (u *nonFormalEducationUsecase) Create(formalEducation *entity.FormalEducation) error {
	return u.formalEducationRepository.Create(formalEducation)
}

func (u *nonFormalEducationUsecase) Update(formalEducation *entity.FormalEducation) error {
	return u.formalEducationRepository.Update(formalEducation)
}

func (u *nonFormalEducationUsecase) Delete(id uint) error {
  return u.formalEducationRepository.Delete(id)
}

