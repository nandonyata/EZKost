package http

import (
	"ezkost/internal/delivery/http/handler"
	"ezkost/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	authHandler *handler.AuthHandler,
	roomHandler *handler.RoomHandler,
	tenantHandler *handler.TenantHandler,
	paymentHandler *handler.PaymentHandler,
	dashboardHandler *handler.DashboardHandler,
	expenseHandler *handler.ExpenseHandler,
) {
	// API v1
	v1 := r.Group("/api/v1")

	// Public routes
	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	// Protected routes
	protected := v1.Group("")
	protected.Use(authMiddleware.Authenticate())
	{
		// Dashboard
		protected.GET("/dashboard/summary", dashboardHandler.GetSummary)

		// Rooms
		rooms := protected.Group("/rooms")
		{
			rooms.GET("", roomHandler.GetAll)
			rooms.GET("/:id", roomHandler.GetByID)
			rooms.POST("", roomHandler.Create)
			rooms.PUT("/:id", roomHandler.Update)
			rooms.DELETE("/:id", roomHandler.Delete)
		}

		// Tenants
		tenants := protected.Group("/tenants")
		{
			tenants.GET("", tenantHandler.GetAll)
			tenants.GET("/:id", tenantHandler.GetByID)
			tenants.POST("", tenantHandler.Create)
			tenants.PUT("/:id", tenantHandler.Update)
			tenants.DELETE("/:id", tenantHandler.Delete)
		}

		// Payments
		payments := protected.Group("/payments")
		{
			payments.GET("", paymentHandler.GetAll)
			payments.GET("/:id", paymentHandler.GetByID)
			payments.GET("/tenant/:tenant_id", paymentHandler.GetByTenantID)
			payments.GET("/overdue", paymentHandler.GetOverdue)
			payments.POST("", paymentHandler.Create)
			payments.PUT("/:id", paymentHandler.Update)
		}

		// Expenses
		expenses := protected.Group("/expenses")
		{
			expenses.GET("", expenseHandler.GetAll)
			expenses.GET("/:id", expenseHandler.GetByID)
			expenses.POST("", expenseHandler.Create)
			expenses.PUT("/:id", expenseHandler.Update)
			expenses.DELETE("/:id", expenseHandler.Delete)
		}
	}
}
