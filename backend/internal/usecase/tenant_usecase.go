package usecase

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"time"
)

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
