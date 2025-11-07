package service

import (
	"errors"
	"fmt"

	"github.com/mlm-app/backend/internal/domain"
	"github.com/mlm-app/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type DistributorService interface {
	Register(distributor *domain.Distributor, password string) error
	Login(email, password string) (*domain.Distributor, error)
	GetByID(id uint) (*domain.Distributor, error)
	GetByEmail(email string) (*domain.Distributor, error)
	Update(distributor *domain.Distributor) error
	Delete(id uint) error
	List(offset, limit int) ([]domain.Distributor, int64, error)
	GetDownlines(sponsorID uint) ([]domain.Distributor, error)
	GetTreeStructure(distributorID uint, depth int) (*domain.TreeNode, error)
	AddMemberToTree(member *domain.Distributor, sponsorID uint, position string) error
	CheckRankEligibility(distributorID uint) (*domain.Rank, error)
	UpdateRank(distributorID, rankID uint) error
}

type distributorService struct {
	distributorRepo repository.DistributorRepository
	rankRepo        repository.RankRepository
	treeService     TreeService
}

func NewDistributorService(
	distributorRepo repository.DistributorRepository,
	rankRepo repository.RankRepository,
	treeService TreeService,
) DistributorService {
	return &distributorService{
		distributorRepo: distributorRepo,
		rankRepo:        rankRepo,
		treeService:     treeService,
	}
}

// Register creates a new distributor
func (s *distributorService) Register(distributor *domain.Distributor, password string) error {
	// Check if email already exists
	existing, _ := s.distributorRepo.FindByEmail(distributor.Email)
	if existing != nil {
		return errors.New("email already registered")
	}
	
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	distributor.PasswordHash = string(hashedPassword)
	
	// If has sponsor, validate and set level
	if distributor.SponsorID != nil {
		sponsor, err := s.distributorRepo.FindByID(*distributor.SponsorID)
		if err != nil {
			return errors.New("invalid sponsor ID")
		}
		
		// Calculate level
		level, err := s.treeService.CalculateLevel(*distributor.SponsorID)
		if err != nil {
			return err
		}
		distributor.Level = level
		
		// If no tree type specified, use sponsor's tree type
		if distributor.TreeType == "" {
			distributor.TreeType = sponsor.TreeType
		}
		
		// If no position specified, find available position
		if distributor.Position == "" {
			position, err := s.treeService.FindAvailablePosition(*distributor.SponsorID, distributor.TreeType)
			if err != nil {
				return err
			}
			distributor.Position = position
		} else {
			// Validate the specified position
			if err := s.treeService.ValidatePosition(*distributor.SponsorID, distributor.TreeType, distributor.Position); err != nil {
				return err
			}
		}
	} else {
		// Root distributor
		distributor.Level = 0
		if distributor.TreeType == "" {
			distributor.TreeType = domain.TreeTypeBinary
		}
	}
	
	// Set default status
	if distributor.Status == "" {
		distributor.Status = "active"
	}
	
	return s.distributorRepo.Create(distributor)
}

// Login authenticates a distributor
func (s *distributorService) Login(email, password string) (*domain.Distributor, error) {
	distributor, err := s.distributorRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(distributor.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	
	// Check if account is active
	if distributor.Status != "active" {
		return nil, errors.New("account is not active")
	}
	
	return distributor, nil
}

// GetByID retrieves a distributor by ID
func (s *distributorService) GetByID(id uint) (*domain.Distributor, error) {
	return s.distributorRepo.FindByID(id)
}

// GetByEmail retrieves a distributor by email
func (s *distributorService) GetByEmail(email string) (*domain.Distributor, error) {
	return s.distributorRepo.FindByEmail(email)
}

// Update updates a distributor
func (s *distributorService) Update(distributor *domain.Distributor) error {
	return s.distributorRepo.Update(distributor)
}

// Delete deletes a distributor
func (s *distributorService) Delete(id uint) error {
	return s.distributorRepo.Delete(id)
}

// List retrieves a list of distributors
func (s *distributorService) List(offset, limit int) ([]domain.Distributor, int64, error) {
	return s.distributorRepo.List(offset, limit)
}

// GetDownlines retrieves all downlines for a distributor
func (s *distributorService) GetDownlines(sponsorID uint) ([]domain.Distributor, error) {
	return s.distributorRepo.GetDownlines(sponsorID)
}

// GetTreeStructure retrieves the tree structure
func (s *distributorService) GetTreeStructure(distributorID uint, depth int) (*domain.TreeNode, error) {
	return s.treeService.GetTreeStructure(distributorID, depth)
}

// AddMemberToTree adds a new member to the tree
func (s *distributorService) AddMemberToTree(member *domain.Distributor, sponsorID uint, position string) error {
	member.SponsorID = &sponsorID
	member.Position = position
	
	// Calculate level
	level, err := s.treeService.CalculateLevel(sponsorID)
	if err != nil {
		return err
	}
	member.Level = level
	
	// Validate position
	sponsor, err := s.distributorRepo.FindByID(sponsorID)
	if err != nil {
		return err
	}
	
	if err := s.treeService.ValidatePosition(sponsorID, sponsor.TreeType, position); err != nil {
		return err
	}
	
	member.TreeType = sponsor.TreeType
	
	return s.distributorRepo.Create(member)
}

// CheckRankEligibility checks if a distributor is eligible for a rank upgrade
func (s *distributorService) CheckRankEligibility(distributorID uint) (*domain.Rank, error) {
	distributor, err := s.distributorRepo.FindByID(distributorID)
	if err != nil {
		return nil, err
	}
	
	// Get all ranks ordered by level
	ranks, err := s.rankRepo.List()
	if err != nil {
		return nil, err
	}
	
	// Find the highest rank the distributor qualifies for
	var eligibleRank *domain.Rank
	currentRankLevel := 0
	if distributor.Rank != nil {
		currentRankLevel = distributor.Rank.Level
	}
	
	for _, rank := range ranks {
		// Skip if already at this rank or lower
		if rank.Level <= currentRankLevel {
			continue
		}
		
		// Check requirements
		if s.meetsRankRequirements(distributor, &rank) {
			eligibleRank = &rank
		}
	}
	
	return eligibleRank, nil
}

// meetsRankRequirements checks if distributor meets rank requirements
func (s *distributorService) meetsRankRequirements(distributor *domain.Distributor, rank *domain.Rank) bool {
	// Check personal sales
	if distributor.PersonalSales < rank.MinPersonalSales {
		return false
	}
	
	// Check team sales
	if distributor.TeamSales < rank.MinTeamSales {
		return false
	}
	
	// Check downlines count
	downlineCount, _ := s.distributorRepo.CountDownlines(distributor.ID)
	if int(downlineCount) < rank.MinDownlines {
		return false
	}
	
	// Check active downlines count
	activeDownlineCount, _ := s.distributorRepo.CountActiveDownlines(distributor.ID)
	if int(activeDownlineCount) < rank.MinActiveDownlines {
		return false
	}
	
	return true
}

// UpdateRank updates a distributor's rank
func (s *distributorService) UpdateRank(distributorID, rankID uint) error {
	distributor, err := s.distributorRepo.FindByID(distributorID)
	if err != nil {
		return err
	}
	
	rank, err := s.rankRepo.FindByID(rankID)
	if err != nil {
		return err
	}
	
	// Verify eligibility
	if !s.meetsRankRequirements(distributor, rank) {
		return fmt.Errorf("distributor does not meet requirements for rank: %s", rank.Name)
	}
	
	distributor.RankID = &rankID
	return s.distributorRepo.Update(distributor)
}
