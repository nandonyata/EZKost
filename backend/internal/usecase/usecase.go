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

// Tenant Usecase
type TenantUsecase interface {
	Create(tenant *entity.Tenant) error
	GetAll() ([]entity.Tenant, error)
	GetByID(id uint) (*entity.Tenant, error)
	Update(oldRoomID *uint, tenant *entity.Tenant) error
	Delete(id uint) error
}

type tenantUsecase struct {
	tenantRepo repository.TenantRepository
	roomRepo   repository.RoomRepository
}

func NewTenantUsecase(tenantRepo repository.TenantRepository, roomRepo repository.RoomRepository) TenantUsecase {
	return &tenantUsecase{
		tenantRepo: tenantRepo,
		roomRepo:   roomRepo,
	}
}

func (u *tenantUsecase) Create(tenant *entity.Tenant) error {
	tenant.CreatedAt = time.Now()
	tenant.UpdatedAt = time.Now()

	// Update room status to occupied
	if tenant.RoomID != nil {
		if err := u.roomRepo.UpdateStatus(*tenant.RoomID, "occupied"); err != nil {
			return err
		}
	}

	return u.tenantRepo.Create(tenant)
}

func (u *tenantUsecase) GetAll() ([]entity.Tenant, error) {
	return u.tenantRepo.FindAll()
}

func (u *tenantUsecase) GetByID(id uint) (*entity.Tenant, error) {
	return u.tenantRepo.FindByID(id)
}

func (u *tenantUsecase) Update(oldRoomID *uint, tenant *entity.Tenant) error {
	tenant.UpdatedAt = time.Now()

	// Update old room to empty
	if oldRoomID != nil {
		if err := u.roomRepo.UpdateStatus(*oldRoomID, "empty"); err != nil {
			return err
		}
	}

	// Update new room to occupied
	if tenant.RoomID != nil {
		if err := u.roomRepo.UpdateStatus(*tenant.RoomID, "occupied"); err != nil {
			return err
		}
	}

	return u.tenantRepo.Update(tenant)
}

func (u *tenantUsecase) Delete(id uint) error {
	tenant, err := u.tenantRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Update room to empty
	if tenant.RoomID != nil {
		if err := u.roomRepo.UpdateStatus(*tenant.RoomID, "empty"); err != nil {
			return err
		}
	}

	return u.tenantRepo.Delete(id)
}

// Payment Usecase
type PaymentUsecase interface {
	Create(payment *entity.Payment) error
	GetAll() ([]entity.Payment, error)
	GetByID(id uint) (*entity.Payment, error)
	GetByTenantID(tenantID uint) ([]entity.Payment, error)
	GetOverdue() ([]entity.Payment, error)
	Update(payment *entity.Payment) error
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
	tenantRepo  repository.TenantRepository
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, tenantRepo repository.TenantRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		tenantRepo:  tenantRepo,
	}
}

func (u *paymentUsecase) Create(payment *entity.Payment) error {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()
	return u.paymentRepo.Create(payment)
}

func (u *paymentUsecase) GetAll() ([]entity.Payment, error) {
	return u.paymentRepo.FindAll()
}

func (u *paymentUsecase) GetByID(id uint) (*entity.Payment, error) {
	return u.paymentRepo.FindByID(id)
}

func (u *paymentUsecase) GetByTenantID(tenantID uint) ([]entity.Payment, error) {
	return u.paymentRepo.FindByTenantID(tenantID)
}

func (u *paymentUsecase) GetOverdue() ([]entity.Payment, error) {
	return u.paymentRepo.FindOverdue(time.Now())
}

func (u *paymentUsecase) Update(payment *entity.Payment) error {
	payment.UpdatedAt = time.Now()
	return u.paymentRepo.Update(payment)
}

// Dashboard Usecase
type DashboardSummary struct {
	TotalRooms     int64   `json:"total_rooms"`
	OccupiedRooms  int64   `json:"occupied_rooms"`
	EmptyRooms     int64   `json:"empty_rooms"`
	MonthlyIncome  float64 `json:"monthly_income"`
	MonthlyExpense float64 `json:"monthly_expense"`
	Profit         float64 `json:"profit"`
	OverdueTenants int64   `json:"overdue_tenants"`
	ActiveTenants  int64   `json:"active_tenants"`
}

type DashboardUsecase interface {
	GetSummary() (*DashboardSummary, error)
}

type dashboardUsecase struct {
	roomRepo    repository.RoomRepository
	tenantRepo  repository.TenantRepository
	paymentRepo repository.PaymentRepository
	expenseRepo repository.ExpenseRepository
}

func NewDashboardUsecase(
	roomRepo repository.RoomRepository,
	tenantRepo repository.TenantRepository,
	paymentRepo repository.PaymentRepository,
	expenseRepo repository.ExpenseRepository,
) DashboardUsecase {
	return &dashboardUsecase{
		roomRepo:    roomRepo,
		tenantRepo:  tenantRepo,
		paymentRepo: paymentRepo,
		expenseRepo: expenseRepo,
	}
}

func (u *dashboardUsecase) GetSummary() (*DashboardSummary, error) {
	summary := &DashboardSummary{}

	// Total rooms
	total, err := u.roomRepo.Count()
	if err != nil {
		return nil, err
	}
	summary.TotalRooms = total

	// Occupied rooms
	occupied, err := u.roomRepo.CountByStatus("occupied")
	if err != nil {
		return nil, err
	}
	summary.OccupiedRooms = occupied
	summary.EmptyRooms = total - occupied

	// Monthly income
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	income, err := u.paymentRepo.SumPaidByPeriod(startOfMonth, endOfMonth)
	if err != nil {
		return nil, err
	}
	summary.MonthlyIncome = income

	// Monthly expense
	expense, err := u.expenseRepo.SumByPeriod(startOfMonth, endOfMonth)
	if err != nil {
		return nil, err
	}
	summary.MonthlyExpense = expense
	summary.Profit = income - expense

	// Overdue tenants
	overdue, err := u.paymentRepo.CountOverdue(now)
	if err != nil {
		return nil, err
	}
	summary.OverdueTenants = overdue

	// Active tenants
	active, err := u.tenantRepo.CountByStatus("active")
	if err != nil {
		return nil, err
	}
	summary.ActiveTenants = active

	return summary, nil
}

// Expense Usecase
type ExpenseUsecase interface {
	Create(expense *entity.Expense) error
	GetAll() ([]entity.Expense, error)
	GetByID(id uint) (*entity.Expense, error)
	Update(expense *entity.Expense) error
	Delete(id uint) error
}

type expenseUsecase struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseUsecase(expenseRepo repository.ExpenseRepository) ExpenseUsecase {
	return &expenseUsecase{expenseRepo: expenseRepo}
}

func (u *expenseUsecase) Create(expense *entity.Expense) error {
	expense.CreatedAt = time.Now()
	expense.UpdatedAt = time.Now()
	return u.expenseRepo.Create(expense)
}

func (u *expenseUsecase) GetAll() ([]entity.Expense, error) {
	return u.expenseRepo.FindAll()
}

func (u *expenseUsecase) GetByID(id uint) (*entity.Expense, error) {
	return u.expenseRepo.FindByID(id)
}

func (u *expenseUsecase) Update(expense *entity.Expense) error {
	expense.UpdatedAt = time.Now()
	return u.expenseRepo.Update(expense)
}

func (u *expenseUsecase) Delete(id uint) error {
	return u.expenseRepo.Delete(id)
}
