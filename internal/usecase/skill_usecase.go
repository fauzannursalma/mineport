package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)

type SkillUsecase interface {
	GetAll() ([]entity.Skill, error)
	GetByID(id uint) (entity.Skill, error)
	Create(skill *entity.Skill) error
	Update(skill *entity.Skill) error
	Delete(id uint) error
}

type skillUsecase struct {
	skillRepository repository.SkillRepository
}

func NewSkillUsecase(skillRepository repository.SkillRepository) *skillUsecase {
	return &skillUsecase{skillRepository}
}

func (u *skillUsecase) GetAll() ([]entity.Skill, error) {
  return u.skillRepository.GetAll()
}

func (u *skillUsecase) GetByID(id uint) (entity.Skill, error) {
  return u.skillRepository.GetByID(id)
}

func (u *skillUsecase) Create(skill *entity.Skill) error {
	return u.skillRepository.Create(skill)
}

func (u *skillUsecase) Update(skill *entity.Skill) error {
	return u.skillRepository.Update(skill)
}

func (u *skillUsecase) Delete(id uint) error {
  return u.skillRepository.Delete(id)
}
