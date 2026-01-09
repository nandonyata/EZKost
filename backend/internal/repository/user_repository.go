package repository

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"ezkost/internal/repository/model"

	"gorm.io/gorm"
)

// User Repository Implementation
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	m := &model.User{}
	m.FromEntity(user)
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*user = *m.ToEntity()
	return nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var m model.User
	if err := r.db.Where("email = ?", email).First(&m).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var m model.User
	if err := r.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *userRepository) Update(user *entity.User) error {
	m := &model.User{}
	m.FromEntity(user)
	return r.db.Save(m).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
