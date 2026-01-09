package repository

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"ezkost/internal/repository/model"
	"time"

	"gorm.io/gorm"
)

// Payment Repository Implementation
type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) repository.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *entity.Payment) error {
	m := &model.Payment{}
	m.FromEntity(payment)
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*payment = *m.ToEntity()
	return nil
}

func (r *paymentRepository) FindAll() ([]entity.Payment, error) {
	var models []model.Payment
	if err := r.db.Preload("Tenant.Room").Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Payment, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *paymentRepository) FindByID(id uint) (*entity.Payment, error) {
	var m model.Payment
	if err := r.db.Preload("Tenant.Room").First(&m, id).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *paymentRepository) FindByTenantID(tenantID uint) ([]entity.Payment, error) {
	var models []model.Payment
	if err := r.db.Where("tenant_id = ?", tenantID).Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Payment, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *paymentRepository) FindOverdue(now time.Time) ([]entity.Payment, error) {
	var models []model.Payment
	if err := r.db.Preload("Tenant.Room").
		Where("status = ? AND due_date < ?", "unpaid", now).
		Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Payment, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *paymentRepository) Update(payment *entity.Payment) error {
	m := &model.Payment{}
	m.FromEntity(payment)
	return r.db.Save(m).Error
}

func (r *paymentRepository) CountOverdue(now time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&model.Payment{}).
		Where("status = ? AND due_date < ?", "unpaid", now).
		Count(&count).Error
	return count, err
}

func (r *paymentRepository) SumPaidByPeriod(start, end time.Time) (float64, error) {
	var result struct {
		Total float64
	}
	err := r.db.Model(&model.Payment{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("status = ? AND paid_at >= ? AND paid_at < ?", "paid", start, end).
		Scan(&result).Error
	return result.Total, err
}
