package repository

import (
	"errors"

	"github.com/mlm-app/backend/internal/domain"
	"gorm.io/gorm"
)

type CommissionRepository interface {
	Create(commission *domain.Commission) error
	FindByID(id uint) (*domain.Commission, error)
	Update(commission *domain.Commission) error
	List(offset, limit int) ([]domain.Commission, int64, error)
	ListByDistributor(distributorID uint, offset, limit int) ([]domain.Commission, int64, error)
	GetTotalCommissionByDistributor(distributorID uint) (float64, error)
	GetPendingCommissions(distributorID uint) ([]domain.Commission, error)
	BulkCreate(commissions []domain.Commission) error
}

type commissionRepository struct {
	db *gorm.DB
}

func NewCommissionRepository(db *gorm.DB) CommissionRepository {
	return &commissionRepository{db: db}
}

func (r *commissionRepository) Create(commission *domain.Commission) error {
	return r.db.Create(commission).Error
}

func (r *commissionRepository) FindByID(id uint) (*domain.Commission, error) {
	var commission domain.Commission
	err := r.db.Preload("Distributor").
		Preload("Order").
		Preload("FromDistributor").
		First(&commission, id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("commission not found")
		}
		return nil, err
	}
	return &commission, nil
}

func (r *commissionRepository) Update(commission *domain.Commission) error {
	return r.db.Save(commission).Error
}

func (r *commissionRepository) List(offset, limit int) ([]domain.Commission, int64, error) {
	var commissions []domain.Commission
	var total int64
	
	err := r.db.Model(&domain.Commission{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = r.db.Preload("Distributor").
		Preload("Order").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&commissions).Error
	
	return commissions, total, err
}

func (r *commissionRepository) ListByDistributor(distributorID uint, offset, limit int) ([]domain.Commission, int64, error) {
	var commissions []domain.Commission
	var total int64
	
	query := r.db.Model(&domain.Commission{}).Where("distributor_id = ?", distributorID)
	
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = query.Preload("Order").
		Preload("FromDistributor").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&commissions).Error
	
	return commissions, total, err
}

func (r *commissionRepository) GetTotalCommissionByDistributor(distributorID uint) (float64, error) {
	var total float64
	err := r.db.Model(&domain.Commission{}).
		Where("distributor_id = ? AND status = ?", distributorID, "paid").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&total).Error
	return total, err
}

func (r *commissionRepository) GetPendingCommissions(distributorID uint) ([]domain.Commission, error) {
	var commissions []domain.Commission
	err := r.db.Where("distributor_id = ? AND status = ?", distributorID, "pending").
		Preload("Order").
		Find(&commissions).Error
	return commissions, err
}

func (r *commissionRepository) BulkCreate(commissions []domain.Commission) error {
	return r.db.Create(&commissions).Error
}
