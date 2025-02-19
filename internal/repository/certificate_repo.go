package repository

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"gorm.io/gorm"
)

type CertificateRepository interface {
	GetAll() ([]entity.Certificate, error)
	GetByID(id uint) (entity.Certificate, error)
	Create(certificate *entity.Certificate) error
	Update(certificate *entity.Certificate) error
	Delete(id uint) error
}

type certificateRepository struct {
  db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) *certificateRepository {
	return &certificateRepository{db}
}

func (r *certificateRepository) GetAll() ([]entity.Certificate, error) {
  var certificates []entity.Certificate
  return certificates, r.db.Find(&certificates).Error
}

func (r *certificateRepository) GetByID(id uint) (entity.Certificate, error) {
  var certificate entity.Certificate
  return certificate, r.db.Where("id =?", id).First(&certificate).Error
}

func (r *certificateRepository) Create(certificate *entity.Certificate) error {
	return r.db.Create(certificate).Error
}

func (r *certificateRepository) Update(certificate *entity.Certificate) error {
	return r.db.Save(certificate).Error
}

func (r *certificateRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Certificate{}, id).Error
}


