package usecase

import (
	"ezkost/internal/domain/repository"
	"time"
)

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
