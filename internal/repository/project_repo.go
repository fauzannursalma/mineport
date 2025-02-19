package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type ProjectRepository interface {
  GetAll() ([]entity.Project, error)
  GetByID(id uint) (entity.Project, error)
  Create(project *entity.Project) error
  Update(project *entity.Project) error
  Delete(id uint) error
}

type projectRepository struct {
  db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *projectRepository {
  return &projectRepository{db}
}

func (r *projectRepository) GetAll() ([]entity.Project, error) {
  var projects []entity.Project
  return projects, r.db.Find(&projects).Error
}

func (r *projectRepository) GetByID(id uint) (entity.Project, error) {
  var project entity.Project
  return project, r.db.First(&project, id).Error
}

func (r *projectRepository) Create(project *entity.Project) error {
  return r.db.Create(project).Error
}

func (r *projectRepository) Update(project *entity.Project) error {
  return r.db.Save(project).Error
}

func (r *projectRepository) Delete(id uint) error {
  return r.db.Delete(&entity.Project{}, id).Error
}
