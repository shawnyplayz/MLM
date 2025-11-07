package service

import (
	"fmt"
	"time"

	"github.com/mlm-app/backend/internal/config"
	"github.com/mlm-app/backend/internal/domain"
	"github.com/mlm-app/backend/internal/repository"
)

type CommissionService interface {
	CalculateAndCreateCommissions(order *domain.Order) error
	CalculateDirectReferralCommission(order *domain.Order) (*domain.Commission, error)
	CalculateLevelCommissions(order *domain.Order) ([]domain.Commission, error)
	CalculateRankBonus(distributorID uint) (*domain.Commission, error)
	ApproveCommission(commissionID uint) error
	PayCommission(commissionID uint) error
	GetDistributorCommissions(distributorID uint, offset, limit int) ([]domain.Commission, int64, error)
}

type commissionService struct {
	commissionRepo  repository.CommissionRepository
	distributorRepo repository.DistributorRepository
	treeService     TreeService
	config          *config.Config
}

func NewCommissionService(
	commissionRepo repository.CommissionRepository,
	distributorRepo repository.DistributorRepository,
	treeService TreeService,
	cfg *config.Config,
) CommissionService {
	return &commissionService{
		commissionRepo:  commissionRepo,
		distributorRepo: distributorRepo,
		treeService:     treeService,
		config:          cfg,
	}
}

// CalculateAndCreateCommissions calculates all commissions for an order
func (s *commissionService) CalculateAndCreateCommissions(order *domain.Order) error {
	var commissions []domain.Commission
	
	// 1. Direct referral commission
	directComm, err := s.CalculateDirectReferralCommission(order)
	if err == nil && directComm != nil {
		commissions = append(commissions, *directComm)
	}
	
	// 2. Level commissions (upline chain)
	levelComms, err := s.CalculateLevelCommissions(order)
	if err == nil && len(levelComms) > 0 {
		commissions = append(commissions, levelComms...)
	}
	
	// 3. Save all commissions
	if len(commissions) > 0 {
		if err := s.commissionRepo.BulkCreate(commissions); err != nil {
			return err
		}
		
		// Update distributor commission totals
		for _, comm := range commissions {
			s.distributorRepo.UpdateCommission(comm.DistributorID, comm.Amount)
		}
	}
	
	return nil
}

// CalculateDirectReferralCommission calculates commission for direct sponsor
func (s *commissionService) CalculateDirectReferralCommission(order *domain.Order) (*domain.Commission, error) {
	distributor, err := s.distributorRepo.FindByID(order.DistributorID)
	if err != nil {
		return nil, err
	}
	
	// If no sponsor, no direct commission
	if distributor.SponsorID == nil {
		return nil, nil
	}
	
	sponsor, err := s.distributorRepo.FindByID(*distributor.SponsorID)
	if err != nil {
		return nil, err
	}
	
	// Calculate commission based on package or default rate
	commissionRate := s.config.MLM.DirectReferralCommission
	if sponsor.Package != nil {
		commissionRate = sponsor.Package.CommissionRate
	}
	
	// Calculate commissionable value from order items
	commissionableValue := s.calculateCommissionableValue(order)
	amount := commissionableValue * (commissionRate / 100)
	
	commission := &domain.Commission{
		DistributorID:     sponsor.ID,
		OrderID:           &order.ID,
		Type:              "direct",
		Level:             1,
		Amount:            amount,
		Percentage:        commissionRate,
		FromDistributorID: &distributor.ID,
		Status:            "pending",
		Description:       fmt.Sprintf("Direct referral commission from order #%s", order.OrderNumber),
	}
	
	return commission, nil
}

// CalculateLevelCommissions calculates commissions for upline chain
func (s *commissionService) CalculateLevelCommissions(order *domain.Order) ([]domain.Commission, error) {
	var commissions []domain.Commission
	
	distributor, err := s.distributorRepo.FindByID(order.DistributorID)
	if err != nil {
		return nil, err
	}
	
	// Get upline chain
	maxLevels := s.config.MLM.MaxCommissionLevels
	if distributor.Package != nil {
		maxLevels = distributor.Package.MaxLevels
	}
	
	upline, err := s.treeService.GetUplineChain(distributor.ID, maxLevels)
	if err != nil {
		return nil, err
	}
	
	commissionableValue := s.calculateCommissionableValue(order)
	levelPercentage := s.config.MLM.LevelCommissionPercentage
	
	// Calculate commission for each level
	for i, uplineDistributor := range upline {
		level := i + 2 // Level 1 is direct, so start from 2
		
		// Skip if distributor is not active
		if uplineDistributor.Status != "active" {
			continue
		}
		
		// Calculate commission with decreasing percentage
		percentage := levelPercentage / float64(level)
		amount := commissionableValue * (percentage / 100)
		
		commission := domain.Commission{
			DistributorID:     uplineDistributor.ID,
			OrderID:           &order.ID,
			Type:              "level",
			Level:             level,
			Amount:            amount,
			Percentage:        percentage,
			FromDistributorID: &distributor.ID,
			Status:            "pending",
			Description:       fmt.Sprintf("Level %d commission from order #%s", level, order.OrderNumber),
		}
		
		commissions = append(commissions, commission)
	}
	
	return commissions, nil
}

// CalculateRankBonus calculates rank achievement bonus
func (s *commissionService) CalculateRankBonus(distributorID uint) (*domain.Commission, error) {
	distributor, err := s.distributorRepo.FindByID(distributorID)
	if err != nil {
		return nil, err
	}
	
	if distributor.Rank == nil {
		return nil, nil
	}
	
	// Monthly rank bonus
	if distributor.Rank.MonthlyBonus > 0 {
		commission := &domain.Commission{
			DistributorID: distributorID,
			Type:          "rank_bonus",
			Amount:        distributor.Rank.MonthlyBonus,
			Status:        "pending",
			Description:   fmt.Sprintf("Monthly bonus for %s rank", distributor.Rank.Name),
		}
		return commission, nil
	}
	
	return nil, nil
}

// ApproveCommission approves a pending commission
func (s *commissionService) ApproveCommission(commissionID uint) error {
	commission, err := s.commissionRepo.FindByID(commissionID)
	if err != nil {
		return err
	}
	
	if commission.Status != "pending" {
		return fmt.Errorf("commission is not in pending status")
	}
	
	commission.Status = "approved"
	return s.commissionRepo.Update(commission)
}

// PayCommission marks a commission as paid
func (s *commissionService) PayCommission(commissionID uint) error {
	commission, err := s.commissionRepo.FindByID(commissionID)
	if err != nil {
		return err
	}
	
	if commission.Status != "approved" {
		return fmt.Errorf("commission must be approved before payment")
	}
	
	now := time.Now()
	commission.Status = "paid"
	commission.PaidAt = &now
	
	return s.commissionRepo.Update(commission)
}

// GetDistributorCommissions retrieves commissions for a distributor
func (s *commissionService) GetDistributorCommissions(distributorID uint, offset, limit int) ([]domain.Commission, int64, error) {
	return s.commissionRepo.ListByDistributor(distributorID, offset, limit)
}

// calculateCommissionableValue calculates the total commissionable value from order items
func (s *commissionService) calculateCommissionableValue(order *domain.Order) float64 {
	var total float64
	for _, item := range order.OrderItems {
		if item.Product != nil {
			total += item.Product.CommissionableValue * float64(item.Quantity)
		} else {
			// If no specific commissionable value, use the item total
			total += item.Total
		}
	}
	return total
}
