package usecase

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"time"
)

// Room Usecase
type RoomUsecase interface {
	Create(room *entity.Room) error
	GetAll() ([]entity.Room, error)
	GetByID(id uint) (*entity.Room, error)
	Update(room *entity.Room) error
	Delete(id uint) error
}

type roomUsecase struct {
	roomRepo repository.RoomRepository
}

func NewRoomUsecase(roomRepo repository.RoomRepository) RoomUsecase {
	return &roomUsecase{roomRepo: roomRepo}
}

func (u *roomUsecase) Create(room *entity.Room) error {
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()
	return u.roomRepo.Create(room)
}

func (u *roomUsecase) GetAll() ([]entity.Room, error) {
	return u.roomRepo.FindAll()
}

func (u *roomUsecase) GetByID(id uint) (*entity.Room, error) {
	return u.roomRepo.FindByID(id)
}

func (u *roomUsecase) Update(room *entity.Room) error {
	room.UpdatedAt = time.Now()
	return u.roomRepo.Update(room)
}

func (u *roomUsecase) Delete(id uint) error {
	return u.roomRepo.Delete(id)
}
