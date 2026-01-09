package model

import (
	"ezkost/internal/domain/entity"
	"time"
)

type Tenant struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Phone     string    `gorm:"size:20;not null"`
	RoomID    *uint     `gorm:"index"`
	StartDate time.Time `gorm:"not null"`
	EndDate   *time.Time
	Status    string `gorm:"size:20;not null;default:'active'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      *Room     `gorm:"foreignKey:RoomID"`
	Payments  []Payment `gorm:"foreignKey:TenantID"`
}

func (Tenant) TableName() string {
	return "tenants"
}

func (m *Tenant) ToEntity() *entity.Tenant {
	tenant := &entity.Tenant{
		ID:        m.ID,
		Name:      m.Name,
		Phone:     m.Phone,
		RoomID:    m.RoomID,
		StartDate: m.StartDate,
		EndDate:   m.EndDate,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
	if m.Room != nil {
		tenant.Room = m.Room.ToEntity()
	}
	if m.Payments != nil {
		tenant.Payments = make([]entity.Payment, len(m.Payments))
		for i, p := range m.Payments {
			tenant.Payments[i] = *p.ToEntity()
		}
	}
	return tenant
}

func (m *Tenant) FromEntity(e *entity.Tenant) {
	m.ID = e.ID
	m.Name = e.Name
	m.Phone = e.Phone
	m.RoomID = e.RoomID
	m.StartDate = e.StartDate
	m.EndDate = e.EndDate
	m.Status = e.Status
}
