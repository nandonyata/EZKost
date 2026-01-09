package repository

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"ezkost/internal/repository/model"
	"time"

	"gorm.io/gorm"
)

// Expense Repository Implementation
type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) repository.ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) Create(expense *entity.Expense) error {
	m := &model.Expense{}
	m.FromEntity(expense)
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*expense = *m.ToEntity()
	return nil
}

func (r *expenseRepository) FindAll() ([]entity.Expense, error) {
	var models []model.Expense
	if err := r.db.Order("expense_date DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Expense, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *expenseRepository) FindByID(id uint) (*entity.Expense, error) {
	var m model.Expense
	if err := r.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *expenseRepository) Update(expense *entity.Expense) error {
	m := &model.Expense{}
	m.FromEntity(expense)
	return r.db.Save(m).Error
}

func (r *expenseRepository) Delete(id uint) error {
	return r.db.Delete(&model.Expense{}, id).Error
}

func (r *expenseRepository) SumByPeriod(start, end time.Time) (float64, error) {
	var result struct {
		Total float64
	}
	err := r.db.Model(&model.Expense{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("expense_date >= ? AND expense_date < ?", start, end).
		Scan(&result).Error
	return result.Total, err
}
