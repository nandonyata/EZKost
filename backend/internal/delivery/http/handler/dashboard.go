package handler

import (
	"ezkost/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dashboard Handler
type DashboardHandler struct {
	dashboardUsecase usecase.DashboardUsecase
}

func NewDashboardHandler(dashboardUsecase usecase.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{dashboardUsecase: dashboardUsecase}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	summary, err := h.dashboardUsecase.GetSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
