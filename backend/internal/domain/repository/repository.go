package repository

import (
	"ezkost/internal/domain/entity"
	"time"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type RoomRepository interface {
	Create(room *entity.Room) error
	FindAll() ([]entity.Room, error)
	FindByID(id uint) (*entity.Room, error)
	Update(room *entity.Room) error
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
	Count() (int64, error)
	CountByStatus(status string) (int64, error)
}

type TenantRepository interface {
	Create(tenant *entity.Tenant) error
	FindAll() ([]entity.Tenant, error)
	FindByID(id uint) (*entity.Tenant, error)
	Update(tenant *entity.Tenant) error
	Delete(id uint) error
	CountByStatus(status string) (int64, error)
}

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

type ExpenseRepository interface {
	Create(expense *entity.Expense) error
	FindAll() ([]entity.Expense, error)
	FindByID(id uint) (*entity.Expense, error)
	Update(expense *entity.Expense) error
	Delete(id uint) error
	SumByPeriod(start, end time.Time) (float64, error)
}
