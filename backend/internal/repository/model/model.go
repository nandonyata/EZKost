package model

import (
	"ezkost/internal/domain/entity"
	"time"

	"gorm.io/gorm"
)

// GORM Models (Database Layer)
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Email        string `gorm:"size:100;unique;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	Role         string `gorm:"size:20;not null;default:'staff'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (User) TableName() string {
	return "users"
}

func (m *User) ToEntity() *entity.User {
	return &entity.User{
		ID:           m.ID,
		Name:         m.Name,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		Role:         m.Role,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (m *User) FromEntity(e *entity.User) {
	m.ID = e.ID
	m.Name = e.Name
	m.Email = e.Email
	m.PasswordHash = e.PasswordHash
	m.Role = e.Role
}

type Room struct {
	ID         uint    `gorm:"primaryKey"`
	RoomNumber string  `gorm:"size:20;unique;not null"`
	Price      float64 `gorm:"not null"`
	Status     string  `gorm:"size:20;not null;default:'empty'"`
	Facilities string  `gorm:"type:text"`
	Notes      string  `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Tenant     *Tenant `gorm:"foreignKey:RoomID"`
}

func (Room) TableName() string {
	return "rooms"
}

func (m *Room) ToEntity() *entity.Room {
	room := &entity.Room{
		ID:         m.ID,
		RoomNumber: m.RoomNumber,
		Price:      m.Price,
		Status:     m.Status,
		Facilities: m.Facilities,
		Notes:      m.Notes,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
	if m.Tenant != nil {
		room.Tenant = m.Tenant.ToEntity()
	}
	return room
}

func (m *Room) FromEntity(e *entity.Room) {
	m.ID = e.ID
	m.RoomNumber = e.RoomNumber
	m.Price = e.Price
	m.Status = e.Status
	m.Facilities = e.Facilities
	m.Notes = e.Notes
}

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
