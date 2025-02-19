package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)

type ProjectUsecase interface {
	GetAll() ([]entity.Project, error)
	GetByID(id uint) (entity.Project, error)
	Create(project *entity.Project) error
	Update(project *entity.Project) error
	Delete(id uint) error
}

type projectUsecase struct {
  projectRepository repository.ProjectRepository
}

func NewProjectUsecase(projectRepository repository.ProjectRepository) *projectUsecase {
  return &projectUsecase{projectRepository}
}

func (u *projectUsecase) GetAll() ([]entity.Project, error) {
  return u.projectRepository.GetAll()
}

func (u *projectUsecase) GetByID(id uint) (entity.Project, error) {
  return u.projectRepository.GetByID(id)
}

func (u *projectUsecase) Create(project *entity.Project) error {
  return u.projectRepository.Create(project)
}

func (u *projectUsecase) Update(project *entity.Project) error {
  return u.projectRepository.Update(project)
}

func (u *projectUsecase) Delete(id uint) error {
  return u.projectRepository.Delete(id)
}