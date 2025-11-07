package repository

import (
	"errors"

	"github.com/mlm-app/backend/internal/domain"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	FindByID(id uint) (*domain.Order, error)
	FindByOrderNumber(orderNumber string) (*domain.Order, error)
	Update(order *domain.Order) error
	List(offset, limit int) ([]domain.Order, int64, error)
	ListByDistributor(distributorID uint, offset, limit int) ([]domain.Order, int64, error)
	GetTotalSalesByDistributor(distributorID uint) (float64, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Distributor").
		Preload("OrderItems.Product").
		First(&order, id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) FindByOrderNumber(orderNumber string) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Where("order_number = ?", orderNumber).
		Preload("Distributor").
		Preload("OrderItems.Product").
		First(&order).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) List(offset, limit int) ([]domain.Order, int64, error) {
	var orders []domain.Order
	var total int64
	
	err := r.db.Model(&domain.Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = r.db.Preload("Distributor").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&orders).Error
	
	return orders, total, err
}

func (r *orderRepository) ListByDistributor(distributorID uint, offset, limit int) ([]domain.Order, int64, error) {
	var orders []domain.Order
	var total int64
	
	query := r.db.Model(&domain.Order{}).Where("distributor_id = ?", distributorID)
	
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = query.Preload("OrderItems.Product").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&orders).Error
	
	return orders, total, err
}

func (r *orderRepository) GetTotalSalesByDistributor(distributorID uint) (float64, error) {
	var total float64
	err := r.db.Model(&domain.Order{}).
		Where("distributor_id = ? AND status = ?", distributorID, "completed").
		Select("COALESCE(SUM(total), 0)").
		Scan(&total).Error
	return total, err
}
