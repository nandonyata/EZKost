package model

import (
	"ezkost/internal/domain/entity"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	TenantID      uint      `gorm:"not null;index"`
	Amount        float64   `gorm:"not null"`
	DueDate       time.Time `gorm:"not null"`
	PaidAt        *time.Time
	Status        string `gorm:"size:20;not null;default:'unpaid'"`
	PaymentMethod string `gorm:"size:20"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Tenant        Tenant `gorm:"foreignKey:TenantID"`
}

func (Payment) TableName() string {
	return "payments"
}

func (m *Payment) ToEntity() *entity.Payment {
	return &entity.Payment{
		ID:            m.ID,
		TenantID:      m.TenantID,
		Amount:        m.Amount,
		DueDate:       m.DueDate,
		PaidAt:        m.PaidAt,
		Status:        m.Status,
		PaymentMethod: m.PaymentMethod,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		Tenant:        *m.Tenant.ToEntity(),
	}
}

func (m *Payment) FromEntity(e *entity.Payment) {
	m.ID = e.ID
	m.TenantID = e.TenantID
	m.Amount = e.Amount
	m.DueDate = e.DueDate
	m.PaidAt = e.PaidAt
	m.Status = e.Status
	m.PaymentMethod = e.PaymentMethod
}

func (p *Payment) BeforeUpdate(tx *gorm.DB) error {
	if p.PaidAt != nil {
		now := time.Now()
		if p.DueDate.Before(now) {
			p.Status = "late"
		} else {
			p.Status = "paid"
		}
	}
	return nil
}
