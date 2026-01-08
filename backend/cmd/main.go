package main

import (
	"ezkost/internal/config"
	"ezkost/internal/delivery/http"
	"ezkost/internal/delivery/http/handler"
	"ezkost/internal/delivery/http/middleware"
	"ezkost/internal/repository"
	"ezkost/internal/usecase"
	"ezkost/package/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db := database.ConnectDB(cfg)

	// Auto migrate
	database.AutoMigrate(db)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	tenantRepo := repository.NewTenantRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)

	// Initialize use cases
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret)
	roomUsecase := usecase.NewRoomUsecase(roomRepo)
	tenantUsecase := usecase.NewTenantUsecase(tenantRepo, roomRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, tenantRepo)
	dashboardUsecase := usecase.NewDashboardUsecase(roomRepo, tenantRepo, paymentRepo, expenseRepo)
	expenseUsecase := usecase.NewExpenseUsecase(expenseRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	roomHandler := handler.NewRoomHandler(roomUsecase)
	tenantHandler := handler.NewTenantHandler(tenantUsecase)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase)
	dashboardHandler := handler.NewDashboardHandler(dashboardUsecase)
	expenseHandler := handler.NewExpenseHandler(expenseUsecase)

	// Setup Gin
	r := gin.Default()

	// Setup middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	// Setup routes
	http.SetupRoutes(r, authMiddleware, authHandler, roomHandler, tenantHandler, paymentHandler, dashboardHandler, expenseHandler)

	// Run server
	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
