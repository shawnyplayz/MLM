package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mlm-app/backend/internal/config"
	"github.com/mlm-app/backend/internal/controller"
	"github.com/mlm-app/backend/internal/middleware"
	"github.com/mlm-app/backend/internal/repository"
	"github.com/mlm-app/backend/internal/service"
	"github.com/mlm-app/backend/pkg/database"
)

func main() {
	// Load configuration
	cfg := config.Load()
	
	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	
	// Run migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	
	// Seed initial data
	if err := database.SeedData(db); err != nil {
		log.Println("Warning: Failed to seed data:", err)
	}
	
	// Initialize repositories
	distributorRepo := repository.NewDistributorRepository(db)
	// orderRepo := repository.NewOrderRepository(db) // TODO: Add order controller
	// commissionRepo := repository.NewCommissionRepository(db) // TODO: Add commission controller
	rankRepo := repository.NewRankRepository(db)
	
	// Initialize services
	treeService := service.NewTreeService(distributorRepo)
	distributorService := service.NewDistributorService(distributorRepo, rankRepo, treeService)
	// commissionService := service.NewCommissionService(commissionRepo, distributorRepo, treeService, cfg) // TODO: Add commission controller
	
	// Initialize controllers
	distributorController := controller.NewDistributorController(distributorService, cfg)
	
	// Setup Gin router
	gin.SetMode(cfg.Server.GinMode)
	router := gin.Default()
	
	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.Origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	
	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		distributors := v1.Group("/distributors")
		{
			distributors.POST("/register", distributorController.Register)
			distributors.POST("/login", distributorController.Login)
		}
		
		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			// Distributor routes
			protected.GET("/distributors/profile", distributorController.GetProfile)
			protected.PUT("/distributors/profile", distributorController.Update)
			protected.GET("/distributors/:id", distributorController.GetByID)
			protected.GET("/distributors", distributorController.List)
			protected.GET("/distributors/:id/downlines", distributorController.GetDownlines)
			protected.GET("/distributors/:id/tree", distributorController.GetTreeStructure)
			protected.POST("/distributors/add-member", distributorController.AddMemberToTree)
		}
	}
	
	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
