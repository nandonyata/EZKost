package repository

import (
	"ezkost/internal/domain/entity"
)

type TenantRepository interface {
	Create(tenant *entity.Tenant) error
	FindAll() ([]entity.Tenant, error)
	FindByID(id uint) (*entity.Tenant, error)
	Update(tenant *entity.Tenant) error
	Delete(id uint) error
	CountByStatus(status string) (int64, error)
}
