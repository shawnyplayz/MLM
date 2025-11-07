package service

import (
	"errors"
	"fmt"

	"github.com/mlm-app/backend/internal/domain"
	"github.com/mlm-app/backend/internal/repository"
)

type TreeService interface {
	FindAvailablePosition(sponsorID uint, treeType domain.TreeType) (string, error)
	ValidatePosition(sponsorID uint, treeType domain.TreeType, position string) error
	GetTreeStructure(distributorID uint, depth int) (*domain.TreeNode, error)
	CalculateLevel(sponsorID uint) (int, error)
	GetUplineChain(distributorID uint, levels int) ([]domain.Distributor, error)
}

type treeService struct {
	distributorRepo repository.DistributorRepository
}

func NewTreeService(distributorRepo repository.DistributorRepository) TreeService {
	return &treeService{
		distributorRepo: distributorRepo,
	}
}

// FindAvailablePosition finds the next available position in the tree
func (s *treeService) FindAvailablePosition(sponsorID uint, treeType domain.TreeType) (string, error) {
	switch treeType {
	case domain.TreeTypeBinary:
		return s.findBinaryPosition(sponsorID)
	case domain.TreeTypeMatrix:
		return s.findMatrixPosition(sponsorID)
	case domain.TreeTypeUnilevel:
		return "direct", nil // Unilevel has no position restrictions
	case domain.TreeTypeBreakaway:
		return s.findBreakawayPosition(sponsorID)
	case domain.TreeTypeHybrid:
		return s.findHybridPosition(sponsorID)
	default:
		return "", errors.New("invalid tree type")
	}
}

// findBinaryPosition finds available position in binary tree (left or right)
func (s *treeService) findBinaryPosition(sponsorID uint) (string, error) {
	// Check if left position is available
	leftDist, err := s.distributorRepo.GetByTreeTypeAndPosition(sponsorID, domain.TreeTypeBinary, "left")
	if err != nil {
		return "", err
	}
	if leftDist == nil {
		return "left", nil
	}
	
	// Check if right position is available
	rightDist, err := s.distributorRepo.GetByTreeTypeAndPosition(sponsorID, domain.TreeTypeBinary, "right")
	if err != nil {
		return "", err
	}
	if rightDist == nil {
		return "right", nil
	}
	
	// Both positions filled, need to find position in downline (spillover)
	// Try left leg first
	position, err := s.findBinaryPosition(leftDist.ID)
	if err == nil {
		return position, nil
	}
	
	// Try right leg
	return s.findBinaryPosition(rightDist.ID)
}

// findMatrixPosition finds available position in matrix tree
func (s *treeService) findMatrixPosition(sponsorID uint) (string, error) {
	// Get all direct downlines
	downlines, err := s.distributorRepo.GetDownlines(sponsorID)
	if err != nil {
		return "", err
	}
	
	// Matrix typically has a width limit (e.g., 3x9 means 3 width)
	// For simplicity, we'll use position as "pos_1", "pos_2", "pos_3"
	matrixWidth := 3 // This should come from config
	
	if len(downlines) < matrixWidth {
		return fmt.Sprintf("pos_%d", len(downlines)+1), nil
	}
	
	// All positions filled at this level, find in downline
	for _, downline := range downlines {
		position, err := s.findMatrixPosition(downline.ID)
		if err == nil {
			return position, nil
		}
	}
	
	return "", errors.New("no available positions in matrix")
}

// findBreakawayPosition finds position in breakaway tree
func (s *treeService) findBreakawayPosition(sponsorID uint) (string, error) {
	// Breakaway is similar to unilevel until breakaway occurs
	downlines, err := s.distributorRepo.GetDownlines(sponsorID)
	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("pos_%d", len(downlines)+1), nil
}

// findHybridPosition finds position in hybrid tree
func (s *treeService) findHybridPosition(sponsorID uint) (string, error) {
	// Hybrid combines multiple tree types
	// For this implementation, we'll use binary as primary
	return s.findBinaryPosition(sponsorID)
}

// ValidatePosition validates if a position is valid for the tree type
func (s *treeService) ValidatePosition(sponsorID uint, treeType domain.TreeType, position string) error {
	switch treeType {
	case domain.TreeTypeBinary:
		if position != "left" && position != "right" {
			return errors.New("binary tree only supports 'left' or 'right' positions")
		}
		// Check if position is already taken
		existing, err := s.distributorRepo.GetByTreeTypeAndPosition(sponsorID, treeType, position)
		if err != nil {
			return err
		}
		if existing != nil {
			return errors.New("position already occupied")
		}
	case domain.TreeTypeMatrix:
		// Validate matrix position format
		// Position should be like "pos_1", "pos_2", etc.
	case domain.TreeTypeUnilevel:
		// Unilevel has no position restrictions
	case domain.TreeTypeBreakaway:
		// Similar to unilevel
	case domain.TreeTypeHybrid:
		// Validate based on primary tree type
	}
	
	return nil
}

// GetTreeStructure retrieves the tree structure starting from a distributor
func (s *treeService) GetTreeStructure(distributorID uint, depth int) (*domain.TreeNode, error) {
	return s.distributorRepo.GetTreeStructure(distributorID, depth)
}

// CalculateLevel calculates the level of a distributor in the tree
func (s *treeService) CalculateLevel(sponsorID uint) (int, error) {
	if sponsorID == 0 {
		return 0, nil
	}
	
	sponsor, err := s.distributorRepo.FindByID(sponsorID)
	if err != nil {
		return 0, err
	}
	
	return sponsor.Level + 1, nil
}

// GetUplineChain retrieves the upline chain for a distributor
func (s *treeService) GetUplineChain(distributorID uint, levels int) ([]domain.Distributor, error) {
	var upline []domain.Distributor
	currentID := distributorID
	
	for i := 0; i < levels; i++ {
		distributor, err := s.distributorRepo.FindByID(currentID)
		if err != nil {
			break
		}
		
		if distributor.SponsorID == nil {
			break
		}
		
		sponsor, err := s.distributorRepo.FindByID(*distributor.SponsorID)
		if err != nil {
			break
		}
		
		upline = append(upline, *sponsor)
		currentID = sponsor.ID
	}
	
	return upline, nil
}
