package model

import (
	"ezkost/internal/domain/entity"
	"time"
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
