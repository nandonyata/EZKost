package entity

import "time"

type User struct {
	ID           uint
	Name         string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Room struct {
	ID         uint
	RoomNumber string
	Price      float64
	Status     string
	Facilities string
	Notes      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Tenant     *Tenant
}

type Tenant struct {
	ID        uint
	Name      string
	Phone     string
	RoomID    *uint
	StartDate time.Time
	EndDate   *time.Time
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      *Room
	Payments  []Payment
}

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

type Expense struct {
	ID          uint
	Description string
	Amount      float64
	ExpenseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
