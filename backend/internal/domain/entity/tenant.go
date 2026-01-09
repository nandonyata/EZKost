package entity

import "time"

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
