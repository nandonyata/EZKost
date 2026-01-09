package repository

import (
	"ezkost/internal/domain/entity"
	"time"
)

type ExpenseRepository interface {
	Create(expense *entity.Expense) error
	FindAll() ([]entity.Expense, error)
	FindByID(id uint) (*entity.Expense, error)
	Update(expense *entity.Expense) error
	Delete(id uint) error
	SumByPeriod(start, end time.Time) (float64, error)
}
