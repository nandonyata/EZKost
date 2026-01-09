package entity

import "time"

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
