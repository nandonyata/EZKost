package repository

import (
	"ezkost/internal/domain/entity"
	"time"
)

type PaymentRepository interface {
	Create(payment *entity.Payment) error
	FindAll() ([]entity.Payment, error)
	FindByID(id uint) (*entity.Payment, error)
	FindByTenantID(tenantID uint) ([]entity.Payment, error)
	FindOverdue(now time.Time) ([]entity.Payment, error)
	Update(payment *entity.Payment) error
	CountOverdue(now time.Time) (int64, error)
	SumPaidByPeriod(start, end time.Time) (float64, error)
}
