package repository

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"ezkost/internal/repository/model"
	"time"

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

// Expense Repository Implementation
type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) repository.ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) Create(expense *entity.Expense) error {
	m := &model.Expense{}
	m.FromEntity(expense)
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	*expense = *m.ToEntity()
	return nil
}

func (r *expenseRepository) FindAll() ([]entity.Expense, error) {
	var models []model.Expense
	if err := r.db.Order("expense_date DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.Expense, len(models))
	for i, m := range models {
		entities[i] = *m.ToEntity()
	}
	return entities, nil
}

func (r *expenseRepository) FindByID(id uint) (*entity.Expense, error) {
	var m model.Expense
	if err := r.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *expenseRepository) Update(expense *entity.Expense) error {
	m := &model.Expense{}
	m.FromEntity(expense)
	return r.db.Save(m).Error
}

func (r *expenseRepository) Delete(id uint) error {
	return r.db.Delete(&model.Expense{}, id).Error
}

func (r *expenseRepository) SumByPeriod(start, end time.Time) (float64, error) {
	var result struct {
		Total float64
	}
	err := r.db.Model(&model.Expense{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("expense_date >= ? AND expense_date < ?", start, end).
		Scan(&result).Error
	return result.Total, err
}
