package model

import (
	"ezkost/internal/domain/entity"
	"time"
)

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
