package handler

import (
	"ezkost/internal/domain/entity"
	"ezkost/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Tenant Handler
type TenantHandler struct {
	tenantUsecase usecase.TenantUsecase
}

func NewTenantHandler(tenantUsecase usecase.TenantUsecase) *TenantHandler {
	return &TenantHandler{tenantUsecase: tenantUsecase}
}

func (h *TenantHandler) GetAll(c *gin.Context) {
	tenants, err := h.tenantUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenants)
}

func (h *TenantHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	tenant, err := h.tenantUsecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}
	c.JSON(http.StatusOK, tenant)
}

func (h *TenantHandler) Create(c *gin.Context) {
	var tenant entity.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.tenantUsecase.Create(&tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tenant)
}

func (h *TenantHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	// Get old tenant data for room tracking
	oldTenant, err := h.tenantUsecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenant not found"})
		return
	}

	var tenant entity.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant.ID = uint(id)
	if err := h.tenantUsecase.Update(oldTenant.RoomID, &tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

func (h *TenantHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.tenantUsecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tenant deleted successfully"})
}
