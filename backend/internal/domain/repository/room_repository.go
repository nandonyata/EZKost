package repository

import (
	"ezkost/internal/domain/entity"
)

type RoomRepository interface {
	Create(room *entity.Room) error
	FindAll() ([]entity.Room, error)
	FindByID(id uint) (*entity.Room, error)
	Update(room *entity.Room) error
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
	Count() (int64, error)
	CountByStatus(status string) (int64, error)
}
