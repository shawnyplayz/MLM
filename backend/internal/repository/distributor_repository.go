package repository

import (
	"errors"

	"github.com/mlm-app/backend/internal/domain"
	"gorm.io/gorm"
)

type DistributorRepository interface {
	Create(distributor *domain.Distributor) error
	FindByID(id uint) (*domain.Distributor, error)
	FindByEmail(email string) (*domain.Distributor, error)
	Update(distributor *domain.Distributor) error
	Delete(id uint) error
	List(offset, limit int) ([]domain.Distributor, int64, error)
	GetDownlines(sponsorID uint) ([]domain.Distributor, error)
	GetDownlinesByLevel(sponsorID uint, level int) ([]domain.Distributor, error)
	GetTreeStructure(distributorID uint, depth int) (*domain.TreeNode, error)
	CountDownlines(sponsorID uint) (int64, error)
	CountActiveDownlines(sponsorID uint) (int64, error)
	GetByTreeTypeAndPosition(sponsorID uint, treeType domain.TreeType, position string) (*domain.Distributor, error)
	UpdateSales(distributorID uint, amount float64) error
	UpdateCommission(distributorID uint, amount float64) error
}

type distributorRepository struct {
	db *gorm.DB
}

func NewDistributorRepository(db *gorm.DB) DistributorRepository {
	return &distributorRepository{db: db}
}

func (r *distributorRepository) Create(distributor *domain.Distributor) error {
	return r.db.Create(distributor).Error
}

func (r *distributorRepository) FindByID(id uint) (*domain.Distributor, error) {
	var distributor domain.Distributor
	err := r.db.Preload("Sponsor").
		Preload("Rank").
		Preload("Package").
		Preload("Downlines").
		First(&distributor, id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("distributor not found")
		}
		return nil, err
	}
	return &distributor, nil
}

func (r *distributorRepository) FindByEmail(email string) (*domain.Distributor, error) {
	var distributor domain.Distributor
	err := r.db.Where("email = ?", email).
		Preload("Sponsor").
		Preload("Rank").
		Preload("Package").
		First(&distributor).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("distributor not found")
		}
		return nil, err
	}
	return &distributor, nil
}

func (r *distributorRepository) Update(distributor *domain.Distributor) error {
	return r.db.Save(distributor).Error
}

func (r *distributorRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Distributor{}, id).Error
}

func (r *distributorRepository) List(offset, limit int) ([]domain.Distributor, int64, error) {
	var distributors []domain.Distributor
	var total int64
	
	err := r.db.Model(&domain.Distributor{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = r.db.Preload("Sponsor").
		Preload("Rank").
		Preload("Package").
		Offset(offset).
		Limit(limit).
		Find(&distributors).Error
	
	return distributors, total, err
}

func (r *distributorRepository) GetDownlines(sponsorID uint) ([]domain.Distributor, error) {
	var downlines []domain.Distributor
	err := r.db.Where("sponsor_id = ?", sponsorID).
		Preload("Rank").
		Preload("Package").
		Find(&downlines).Error
	return downlines, err
}

func (r *distributorRepository) GetDownlinesByLevel(sponsorID uint, level int) ([]domain.Distributor, error) {
	var downlines []domain.Distributor
	err := r.db.Where("sponsor_id = ? AND level = ?", sponsorID, level).
		Preload("Rank").
		Find(&downlines).Error
	return downlines, err
}

func (r *distributorRepository) GetTreeStructure(distributorID uint, depth int) (*domain.TreeNode, error) {
	distributor, err := r.FindByID(distributorID)
	if err != nil {
		return nil, err
	}
	
	node := r.buildTreeNode(distributor)
	
	if depth > 0 {
		r.populateChildren(node, depth-1)
	}
	
	return node, nil
}

func (r *distributorRepository) buildTreeNode(distributor *domain.Distributor) *domain.TreeNode {
	node := &domain.TreeNode{
		ID:            distributor.ID,
		DistributorID: distributor.ID,
		Name:          distributor.FirstName + " " + distributor.LastName,
		Email:         distributor.Email,
		SponsorID:     distributor.SponsorID,
		Position:      distributor.Position,
		Level:         distributor.Level,
		TotalSales:    distributor.TotalSales,
		Status:        distributor.Status,
	}
	
	if distributor.Rank != nil {
		node.RankName = distributor.Rank.Name
	}
	
	return node
}

func (r *distributorRepository) populateChildren(node *domain.TreeNode, depth int) {
	if depth < 0 {
		return
	}
	
	var children []domain.Distributor
	r.db.Where("sponsor_id = ?", node.DistributorID).
		Preload("Rank").
		Find(&children)
	
	for _, child := range children {
		childNode := r.buildTreeNode(&child)
		if depth > 0 {
			r.populateChildren(childNode, depth-1)
		}
		node.Children = append(node.Children, *childNode)
	}
}

func (r *distributorRepository) CountDownlines(sponsorID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Distributor{}).
		Where("sponsor_id = ?", sponsorID).
		Count(&count).Error
	return count, err
}

func (r *distributorRepository) CountActiveDownlines(sponsorID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Distributor{}).
		Where("sponsor_id = ? AND status = ?", sponsorID, "active").
		Count(&count).Error
	return count, err
}

func (r *distributorRepository) GetByTreeTypeAndPosition(sponsorID uint, treeType domain.TreeType, position string) (*domain.Distributor, error) {
	var distributor domain.Distributor
	err := r.db.Where("sponsor_id = ? AND tree_type = ? AND position = ?", sponsorID, treeType, position).
		First(&distributor).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &distributor, nil
}

func (r *distributorRepository) UpdateSales(distributorID uint, amount float64) error {
	return r.db.Model(&domain.Distributor{}).
		Where("id = ?", distributorID).
		UpdateColumn("total_sales", gorm.Expr("total_sales + ?", amount)).
		Error
}

func (r *distributorRepository) UpdateCommission(distributorID uint, amount float64) error {
	return r.db.Model(&domain.Distributor{}).
		Where("id = ?", distributorID).
		UpdateColumn("total_commission", gorm.Expr("total_commission + ?", amount)).
		Error
}
