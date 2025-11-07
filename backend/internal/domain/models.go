package domain

import (
	"time"

	"gorm.io/gorm"
)

// TreeType represents the type of MLM tree structure
type TreeType string

const (
	TreeTypeBinary    TreeType = "binary"
	TreeTypeMatrix    TreeType = "matrix"
	TreeTypeUnilevel  TreeType = "unilevel"
	TreeTypeBreakaway TreeType = "breakaway"
	TreeTypeHybrid    TreeType = "hybrid"
)

// Distributor represents a distributor in the MLM system
type Distributor struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Personal Information
	FirstName         string         `gorm:"size:100;not null" json:"first_name"`
	LastName          string         `gorm:"size:100;not null" json:"last_name"`
	Email             string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Phone             string         `gorm:"size:20" json:"phone"`
	Address           string         `gorm:"size:500" json:"address"`
	City              string         `gorm:"size:100" json:"city"`
	State             string         `gorm:"size:100" json:"state"`
	Country           string         `gorm:"size:100" json:"country"`
	ZipCode           string         `gorm:"size:20" json:"zip_code"`
	DateOfBirth       *time.Time     `json:"date_of_birth"`
	
	// Authentication
	PasswordHash      string         `gorm:"size:255;not null" json:"-"`
	
	// MLM Structure
	SponsorID         *uint          `gorm:"index" json:"sponsor_id"`
	Sponsor           *Distributor   `gorm:"foreignKey:SponsorID" json:"sponsor,omitempty"`
	TreeType          TreeType       `gorm:"size:20;default:'binary'" json:"tree_type"`
	Position          string         `gorm:"size:20" json:"position"` // For binary: left/right
	Level             int            `gorm:"default:0" json:"level"`
	
	// Business Metrics
	TotalSales        float64        `gorm:"type:decimal(15,2);default:0" json:"total_sales"`
	PersonalSales     float64        `gorm:"type:decimal(15,2);default:0" json:"personal_sales"`
	TeamSales         float64        `gorm:"type:decimal(15,2);default:0" json:"team_sales"`
	TotalCommission   float64        `gorm:"type:decimal(15,2);default:0" json:"total_commission"`
	TotalBonus        float64        `gorm:"type:decimal(15,2);default:0" json:"total_bonus"`
	
	// Status and Rank
	Role              string         `gorm:"size:20;default:'distributor'" json:"role"` // admin or distributor
	Status            string         `gorm:"size:20;default:'active'" json:"status"` // active, inactive, suspended
	RankID            *uint          `gorm:"index" json:"rank_id"`
	Rank              *Rank          `gorm:"foreignKey:RankID" json:"rank,omitempty"`
	PackageID         *uint          `gorm:"index" json:"package_id"`
	Package           *Package       `gorm:"foreignKey:PackageID" json:"package,omitempty"`
	
	// Relationships
	Downlines         []Distributor  `gorm:"foreignKey:SponsorID" json:"downlines,omitempty"`
	Orders            []Order        `gorm:"foreignKey:DistributorID" json:"orders,omitempty"`
	Commissions       []Commission   `gorm:"foreignKey:DistributorID" json:"commissions,omitempty"`
	RankAchievements  []RankAchievement `gorm:"foreignKey:DistributorID" json:"rank_achievements,omitempty"`
}

// Rank represents a rank in the MLM system
type Rank struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	
	Name              string         `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Description       string         `gorm:"type:text" json:"description"`
	Level             int            `gorm:"not null;uniqueIndex" json:"level"`
	
	// Requirements
	MinPersonalSales  float64        `gorm:"type:decimal(15,2);default:0" json:"min_personal_sales"`
	MinTeamSales      float64        `gorm:"type:decimal(15,2);default:0" json:"min_team_sales"`
	MinDownlines      int            `gorm:"default:0" json:"min_downlines"`
	MinActiveDownlines int           `gorm:"default:0" json:"min_active_downlines"`
	
	// Benefits
	CommissionBonus   float64        `gorm:"type:decimal(5,2);default:0" json:"commission_bonus"` // Percentage
	MonthlyBonus      float64        `gorm:"type:decimal(15,2);default:0" json:"monthly_bonus"`
	
	Color             string         `gorm:"size:20" json:"color"` // For UI display
	Icon              string         `gorm:"size:100" json:"icon"`
}

// Package represents a distributor package/plan
type Package struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	
	Name              string         `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Description       string         `gorm:"type:text" json:"description"`
	Price             float64        `gorm:"type:decimal(15,2);not null" json:"price"`
	
	// Benefits
	CommissionRate    float64        `gorm:"type:decimal(5,2);default:0" json:"commission_rate"` // Percentage
	MaxLevels         int            `gorm:"default:5" json:"max_levels"`
	
	// Features (JSON stored as text)
	Features          string         `gorm:"type:text" json:"features"`
	
	IsActive          bool           `gorm:"default:true" json:"is_active"`
}

// Order represents a product order
type Order struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	
	OrderNumber       string         `gorm:"size:50;uniqueIndex;not null" json:"order_number"`
	DistributorID     uint           `gorm:"not null;index" json:"distributor_id"`
	Distributor       *Distributor   `gorm:"foreignKey:DistributorID" json:"distributor,omitempty"`
	
	// Order Details
	SubTotal          float64        `gorm:"type:decimal(15,2);not null" json:"sub_total"`
	Tax               float64        `gorm:"type:decimal(15,2);default:0" json:"tax"`
	Shipping          float64        `gorm:"type:decimal(15,2);default:0" json:"shipping"`
	Discount          float64        `gorm:"type:decimal(15,2);default:0" json:"discount"`
	Total             float64        `gorm:"type:decimal(15,2);not null" json:"total"`
	
	Status            string         `gorm:"size:20;default:'pending'" json:"status"` // pending, processing, completed, cancelled
	PaymentStatus     string         `gorm:"size:20;default:'pending'" json:"payment_status"` // pending, paid, failed
	PaymentMethod     string         `gorm:"size:50" json:"payment_method"`
	
	// Shipping Information
	ShippingAddress   string         `gorm:"size:500" json:"shipping_address"`
	ShippingCity      string         `gorm:"size:100" json:"shipping_city"`
	ShippingState     string         `gorm:"size:100" json:"shipping_state"`
	ShippingCountry   string         `gorm:"size:100" json:"shipping_country"`
	ShippingZipCode   string         `gorm:"size:20" json:"shipping_zip_code"`
	
	OrderItems        []OrderItem    `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	
	OrderID           uint           `gorm:"not null;index" json:"order_id"`
	ProductID         uint           `gorm:"not null;index" json:"product_id"`
	Product           *Product       `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	
	Quantity          int            `gorm:"not null" json:"quantity"`
	Price             float64        `gorm:"type:decimal(15,2);not null" json:"price"`
	Total             float64        `gorm:"type:decimal(15,2);not null" json:"total"`
	CommissionableValue float64      `gorm:"type:decimal(15,2);default:0" json:"commissionable_value"`
}

// Product represents a product in the system
type Product struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	
	Name              string         `gorm:"size:255;not null" json:"name"`
	Description       string         `gorm:"type:text" json:"description"`
	SKU               string         `gorm:"size:100;uniqueIndex" json:"sku"`
	
	Price             float64        `gorm:"type:decimal(15,2);not null" json:"price"`
	CostPrice         float64        `gorm:"type:decimal(15,2);default:0" json:"cost_price"`
	CommissionableValue float64      `gorm:"type:decimal(15,2);default:0" json:"commissionable_value"`
	
	Stock             int            `gorm:"default:0" json:"stock"`
	CategoryID        *uint          `gorm:"index" json:"category_id"`
	Category          *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	
	ImageURL          string         `gorm:"size:500" json:"image_url"`
	IsActive          bool           `gorm:"default:true" json:"is_active"`
	IsFeatured        bool           `gorm:"default:false" json:"is_featured"`
}

// Category represents a product category
type Category struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	
	Name              string         `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Description       string         `gorm:"type:text" json:"description"`
	ParentID          *uint          `gorm:"index" json:"parent_id"`
	Parent            *Category      `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	
	IsActive          bool           `gorm:"default:true" json:"is_active"`
}

// Commission represents a commission earned by a distributor
type Commission struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	
	DistributorID     uint           `gorm:"not null;index" json:"distributor_id"`
	Distributor       *Distributor   `gorm:"foreignKey:DistributorID" json:"distributor,omitempty"`
	
	OrderID           *uint          `gorm:"index" json:"order_id"`
	Order             *Order         `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	
	Type              string         `gorm:"size:50;not null" json:"type"` // direct, level, bonus, rank_bonus
	Level             int            `gorm:"default:0" json:"level"`
	Amount            float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Percentage        float64        `gorm:"type:decimal(5,2);default:0" json:"percentage"`
	
	FromDistributorID *uint          `gorm:"index" json:"from_distributor_id"` // Who generated this commission
	FromDistributor   *Distributor   `gorm:"foreignKey:FromDistributorID" json:"from_distributor,omitempty"`
	
	Status            string         `gorm:"size:20;default:'pending'" json:"status"` // pending, approved, paid
	PaidAt            *time.Time     `json:"paid_at"`
	
	Description       string         `gorm:"size:500" json:"description"`
}

// RankAchievement represents when a distributor achieves a rank
type RankAchievement struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	
	DistributorID     uint           `gorm:"not null;index" json:"distributor_id"`
	Distributor       *Distributor   `gorm:"foreignKey:DistributorID" json:"distributor,omitempty"`
	
	RankID            uint           `gorm:"not null;index" json:"rank_id"`
	Rank              *Rank          `gorm:"foreignKey:RankID" json:"rank,omitempty"`
	
	AchievedAt        time.Time      `gorm:"not null" json:"achieved_at"`
	Notes             string         `gorm:"type:text" json:"notes"`
}

// Payout represents a payout to a distributor
type Payout struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	
	DistributorID     uint           `gorm:"not null;index" json:"distributor_id"`
	Distributor       *Distributor   `gorm:"foreignKey:DistributorID" json:"distributor,omitempty"`
	
	Amount            float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Method            string         `gorm:"size:50" json:"method"` // bank_transfer, paypal, check
	
	Status            string         `gorm:"size:20;default:'pending'" json:"status"` // pending, processing, completed, failed
	ProcessedAt       *time.Time     `json:"processed_at"`
	
	TransactionID     string         `gorm:"size:255" json:"transaction_id"`
	Notes             string         `gorm:"type:text" json:"notes"`
}

// TreeNode represents a node in the MLM tree for visualization
type TreeNode struct {
	ID                uint           `json:"id"`
	DistributorID     uint           `json:"distributor_id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	SponsorID         *uint          `json:"sponsor_id"`
	Position          string         `json:"position"`
	Level             int            `json:"level"`
	TotalSales        float64        `json:"total_sales"`
	RankName          string         `json:"rank_name"`
	Status            string         `json:"status"`
	Children          []TreeNode     `json:"children,omitempty"`
}
