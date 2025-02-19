package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)

type WorkExperienceUsecase interface {
	GetAll() ([]entity.WorkExperience, error)
	GetByID(id uint) (entity.WorkExperience, error)
	Create(workExperience *entity.WorkExperience) error
	Update(workExperience *entity.WorkExperience) error
	Delete(id uint) error
}

type workExperienceUsecase struct {
  workExperienceRepository repository.WorkExperienceRepository
}

func NewWorkExperienceUsecase(workExperienceRepository repository.WorkExperienceRepository) *workExperienceUsecase {
  return &workExperienceUsecase{workExperienceRepository}
}

func (u *workExperienceUsecase) GetAll() ([]entity.WorkExperience, error) {
  return u.workExperienceRepository.GetAll()
}

func (u *workExperienceUsecase) GetByID(id uint) (entity.WorkExperience, error) {
  return u.workExperienceRepository.GetByID(id)
}

func (u *workExperienceUsecase) Create(workExperience *entity.WorkExperience) error {
	return u.workExperienceRepository.Create(workExperience)
}

func (u *workExperienceUsecase) Update(workExperience *entity.WorkExperience) error {
  return u.workExperienceRepository.Update(workExperience)
}

func (u *workExperienceUsecase) Delete(id uint) error {
	return u.workExperienceRepository.Delete(id)
}