package database

import (
	"fmt"
	"log"

	"github.com/mlm-app/backend/internal/config"
	"github.com/mlm-app/backend/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	
	log.Println("Database connection established")
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")
	
	err := db.AutoMigrate(
		&domain.Distributor{},
		&domain.Rank{},
		&domain.Package{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Product{},
		&domain.Category{},
		&domain.Commission{},
		&domain.RankAchievement{},
		&domain.Payout{},
	)
	
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	
	log.Println("Database migrations completed")
	return nil
}

func SeedData(db *gorm.DB) error {
	log.Println("Seeding initial data...")
	
	// Seed Ranks
	ranks := []domain.Rank{
		{
			Name:               "Bronze",
			Description:        "Entry level rank",
			Level:              1,
			MinPersonalSales:   0,
			MinTeamSales:       0,
			MinDownlines:       0,
			MinActiveDownlines: 0,
			CommissionBonus:    0,
			MonthlyBonus:       0,
			Color:              "#CD7F32",
			Icon:               "medal",
		},
		{
			Name:               "Silver",
			Description:        "Intermediate rank",
			Level:              2,
			MinPersonalSales:   1000,
			MinTeamSales:       5000,
			MinDownlines:       5,
			MinActiveDownlines: 3,
			CommissionBonus:    2,
			MonthlyBonus:       100,
			Color:              "#C0C0C0",
			Icon:               "medal",
		},
		{
			Name:               "Gold",
			Description:        "Advanced rank",
			Level:              3,
			MinPersonalSales:   5000,
			MinTeamSales:       25000,
			MinDownlines:       15,
			MinActiveDownlines: 10,
			CommissionBonus:    5,
			MonthlyBonus:       500,
			Color:              "#FFD700",
			Icon:               "medal",
		},
		{
			Name:               "Platinum",
			Description:        "Elite rank",
			Level:              4,
			MinPersonalSales:   10000,
			MinTeamSales:       100000,
			MinDownlines:       30,
			MinActiveDownlines: 20,
			CommissionBonus:    10,
			MonthlyBonus:       2000,
			Color:              "#E5E4E2",
			Icon:               "crown",
		},
		{
			Name:               "Diamond",
			Description:        "Top tier rank",
			Level:              5,
			MinPersonalSales:   25000,
			MinTeamSales:       500000,
			MinDownlines:       50,
			MinActiveDownlines: 35,
			CommissionBonus:    15,
			MonthlyBonus:       10000,
			Color:              "#B9F2FF",
			Icon:               "crown",
		},
	}
	
	for _, rank := range ranks {
		var existing domain.Rank
		if err := db.Where("name = ?", rank.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&rank).Error; err != nil {
				return err
			}
		}
	}
	
	// Seed Packages
	packages := []domain.Package{
		{
			Name:           "Starter",
			Description:    "Perfect for beginners",
			Price:          99.99,
			CommissionRate: 10,
			MaxLevels:      5,
			Features:       `["Basic training", "5 level commission", "Email support"]`,
			IsActive:       true,
		},
		{
			Name:           "Professional",
			Description:    "For serious distributors",
			Price:          299.99,
			CommissionRate: 15,
			MaxLevels:      10,
			Features:       `["Advanced training", "10 level commission", "Priority support", "Marketing materials"]`,
			IsActive:       true,
		},
		{
			Name:           "Elite",
			Description:    "Maximum earning potential",
			Price:          999.99,
			CommissionRate: 20,
			MaxLevels:      15,
			Features:       `["Premium training", "15 level commission", "24/7 support", "Marketing materials", "Personal mentor"]`,
			IsActive:       true,
		},
	}
	
	for _, pkg := range packages {
		var existing domain.Package
		if err := db.Where("name = ?", pkg.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&pkg).Error; err != nil {
				return err
			}
		}
	}
	
	// Seed Categories
	categories := []domain.Category{
		{Name: "Health & Wellness", Description: "Health and wellness products", IsActive: true},
		{Name: "Beauty & Personal Care", Description: "Beauty and personal care items", IsActive: true},
		{Name: "Home & Living", Description: "Home and living essentials", IsActive: true},
		{Name: "Nutrition", Description: "Nutritional supplements", IsActive: true},
	}
	
	for _, category := range categories {
		var existing domain.Category
		if err := db.Where("name = ?", category.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&category).Error; err != nil {
				return err
			}
		}
	}
	
	log.Println("Initial data seeded successfully")
	return nil
}
