package entity

import "time"

type Payment struct {
	ID            uint
	TenantID      uint
	Amount        float64
	DueDate       time.Time
	PaidAt        *time.Time
	Status        string
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Tenant        Tenant
}
