package repository

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"ezkost/internal/repository/model"

	"gorm.io/gorm"
)

// Tenant Repository Implementation
type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) repository.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) Create(tenant *entity.Tenant) error {
	m := &model.Tenant{}
	m.FromEntity(tenant)
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*tenant = *m.ToEntity()
	return nil
}

func (r *tenantRepository) FindAll() ([]entity.Tenant, error) {
	var models []model.Tenant
	if err := r.db.Preload("Room").Preload("Payments").Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Tenant, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *tenantRepository) FindByID(id uint) (*entity.Tenant, error) {
	var m model.Tenant
	if err := r.db.Preload("Room").Preload("Payments").First(&m, id).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *tenantRepository) Update(tenant *entity.Tenant) error {
	m := &model.Tenant{}
	m.FromEntity(tenant)
	return r.db.Save(m).Error
}

func (r *tenantRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tenant{}, id).Error
}

func (r *tenantRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Tenant{}).Where("status = ?", status).Count(&count).Error
	return count, err
}
