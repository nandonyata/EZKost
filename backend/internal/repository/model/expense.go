package model

import (
	"ezkost/internal/domain/entity"
	"time"
)

type Expense struct {
	ID          uint      `gorm:"primaryKey"`
	Description string    `gorm:"size:255;not null"`
	Amount      float64   `gorm:"not null"`
	ExpenseDate time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Expense) TableName() string {
	return "expenses"
}

func (m *Expense) ToEntity() *entity.Expense {
	return &entity.Expense{
		ID:          m.ID,
		Description: m.Description,
		Amount:      m.Amount,
		ExpenseDate: m.ExpenseDate,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func (m *Expense) FromEntity(e *entity.Expense) {
	m.ID = e.ID
	m.Description = e.Description
	m.Amount = e.Amount
	m.ExpenseDate = e.ExpenseDate
}
