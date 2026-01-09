package repository

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"ezkost/internal/repository/model"

	"gorm.io/gorm"
)

// Room Repository Implementation
type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) repository.RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) Create(room *entity.Room) error {
	m := &model.Room{}
	m.FromEntity(room)
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*room = *m.ToEntity()
	return nil
}

func (r *roomRepository) FindAll() ([]entity.Room, error) {
	var models []model.Room
	if err := r.db.Preload("Tenant").Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Room, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *roomRepository) FindByID(id uint) (*entity.Room, error) {
	var m model.Room
	if err := r.db.Preload("Tenant").First(&m, id).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *roomRepository) Update(room *entity.Room) error {
	m := &model.Room{}
	m.FromEntity(room)
	return r.db.Save(m).Error
}

func (r *roomRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&model.Room{}).Where("id = ?", id).Update("status", status).Error
}

func (r *roomRepository) Delete(id uint) error {
	return r.db.Delete(&model.Room{}, id).Error
}

func (r *roomRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.Room{}).Count(&count).Error
	return count, err
}

func (r *roomRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Room{}).Where("status = ?", status).Count(&count).Error
	return count, err
}
