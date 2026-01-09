package entity

import "time"

type Expense struct {
	ID          uint
	Description string
	Amount      float64
	ExpenseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
