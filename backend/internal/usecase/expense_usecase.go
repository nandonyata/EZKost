package usecase

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"time"
)

// Expense Usecase
type ExpenseUsecase interface {
	Create(expense *entity.Expense) error
	GetAll() ([]entity.Expense, error)
	GetByID(id uint) (*entity.Expense, error)
	Update(expense *entity.Expense) error
	Delete(id uint) error
}

type expenseUsecase struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseUsecase(expenseRepo repository.ExpenseRepository) ExpenseUsecase {
	return &expenseUsecase{expenseRepo: expenseRepo}
}

func (u *expenseUsecase) Create(expense *entity.Expense) error {
	expense.CreatedAt = time.Now()
	expense.UpdatedAt = time.Now()
	return u.expenseRepo.Create(expense)
}

func (u *expenseUsecase) GetAll() ([]entity.Expense, error) {
	return u.expenseRepo.FindAll()
}

func (u *expenseUsecase) GetByID(id uint) (*entity.Expense, error) {
	return u.expenseRepo.FindByID(id)
}

func (u *expenseUsecase) Update(expense *entity.Expense) error {
	expense.UpdatedAt = time.Now()
	return u.expenseRepo.Update(expense)
}

func (u *expenseUsecase) Delete(id uint) error {
	return u.expenseRepo.Delete(id)
}
