package usecase

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"time"
)

// Payment Usecase
type PaymentUsecase interface {
	Create(payment *entity.Payment) error
	GetAll() ([]entity.Payment, error)
	GetByID(id uint) (*entity.Payment, error)
	GetByTenantID(tenantID uint) ([]entity.Payment, error)
	GetOverdue() ([]entity.Payment, error)
	Update(payment *entity.Payment) error
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
	tenantRepo  repository.TenantRepository
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, tenantRepo repository.TenantRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		tenantRepo:  tenantRepo,
	}
}

func (u *paymentUsecase) Create(payment *entity.Payment) error {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	return u.paymentRepo.Create(payment)
}

func (u *paymentUsecase) GetAll() ([]entity.Payment, error) {
	return u.paymentRepo.FindAll()
}

func (u *paymentUsecase) GetByID(id uint) (*entity.Payment, error) {
	return u.paymentRepo.FindByID(id)
}

func (u *paymentUsecase) GetByTenantID(tenantID uint) ([]entity.Payment, error) {
	return u.paymentRepo.FindByTenantID(tenantID)
}

func (u *paymentUsecase) GetOverdue() ([]entity.Payment, error) {
	return u.paymentRepo.FindOverdue(time.Now())
}

func (u *paymentUsecase) Update(payment *entity.Payment) error {
	payment.UpdatedAt = time.Now()
	return u.paymentRepo.Update(payment)
}
