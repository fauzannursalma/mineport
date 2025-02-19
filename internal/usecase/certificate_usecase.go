package usecase

import (
	"github.com/fauzannursalma/mineport/internal/entity"
	"github.com/fauzannursalma/mineport/internal/repository"
)


type CertificateUsecase interface {
	GetAll() ([]entity.Certificate, error)
	GetByID(id uint) (entity.Certificate, error)
	Create(certificate *entity.Certificate) error
	Update(certificate *entity.Certificate) error
	Delete(id uint) error
}

type certificateUsecase struct {
  certificateRepository repository.CertificateRepository
}

func NewCertificateUsecase(certificateRepository repository.CertificateRepository) *certificateUsecase {
	return &certificateUsecase{certificateRepository}
}

func (u *certificateUsecase) GetAll() ([]entity.Certificate, error) {
  return u.certificateRepository.GetAll()
}

func (u *certificateUsecase) GetByID(id uint) (entity.Certificate, error) {
	return u.certificateRepository.GetByID(id)
}

func (u *certificateUsecase) Create(certificate *entity.Certificate) error {
  return u.certificateRepository.Create(certificate)
}

func (u *certificateUsecase) Update(certificate *entity.Certificate) error {
	return u.certificateRepository.Update(certificate)
}

func (u *certificateUsecase) Delete(id uint) error {
  return u.certificateRepository.Delete(id)
}
