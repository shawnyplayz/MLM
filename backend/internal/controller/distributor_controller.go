package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mlm-app/backend/internal/config"
	"github.com/mlm-app/backend/internal/domain"
	"github.com/mlm-app/backend/internal/middleware"
	"github.com/mlm-app/backend/internal/service"
)

type DistributorController struct {
	distributorService service.DistributorService
	config             *config.Config
}

func NewDistributorController(distributorService service.DistributorService, cfg *config.Config) *DistributorController {
	return &DistributorController{
		distributorService: distributorService,
		config:             cfg,
	}
}

// Register godoc
// @Summary Register a new distributor
// @Tags distributor
// @Accept json
// @Produce json
// @Param distributor body RegisterRequest true "Distributor registration data"
// @Success 201 {object} map[string]interface{}
// @Router /api/v1/distributors/register [post]
func (ctrl *DistributorController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	distributor := &domain.Distributor{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		ZipCode:     req.ZipCode,
		SponsorID:   req.SponsorID,
		TreeType:    req.TreeType,
		Position:    req.Position,
		PackageID:   req.PackageID,
	}
	
	if err := ctrl.distributorService.Register(distributor, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Generate JWT token
	token, err := middleware.GenerateToken(distributor.ID, distributor.Email, distributor.Role, ctrl.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":     "Registration successful",
		"distributor": distributor,
		"token":       token,
	})
}

// Login godoc
// @Summary Login distributor
// @Tags distributor
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/distributors/login [post]
func (ctrl *DistributorController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	distributor, err := ctrl.distributorService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	
	// Generate JWT token
	token, err := middleware.GenerateToken(distributor.ID, distributor.Email, distributor.Role, ctrl.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":     "Login successful",
		"distributor": distributor,
		"token":       token,
	})
}

// GetProfile godoc
// @Summary Get distributor profile
// @Tags distributor
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Distributor
// @Router /api/v1/distributors/profile [get]
func (ctrl *DistributorController) GetProfile(c *gin.Context) {
	distributorID := c.GetUint("distributor_id")
	
	distributor, err := ctrl.distributorService.GetByID(distributorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, distributor)
}

// GetByID godoc
// @Summary Get distributor by ID
// @Tags distributor
// @Produce json
// @Security BearerAuth
// @Param id path int true "Distributor ID"
// @Success 200 {object} domain.Distributor
// @Router /api/v1/distributors/{id} [get]
func (ctrl *DistributorController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	distributor, err := ctrl.distributorService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, distributor)
}

// Update godoc
// @Summary Update distributor
// @Tags distributor
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param distributor body UpdateRequest true "Distributor update data"
// @Success 200 {object} domain.Distributor
// @Router /api/v1/distributors/profile [put]
func (ctrl *DistributorController) Update(c *gin.Context) {
	distributorID := c.GetUint("distributor_id")
	
	distributor, err := ctrl.distributorService.GetByID(distributorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Update fields
	if req.FirstName != "" {
		distributor.FirstName = req.FirstName
	}
	if req.LastName != "" {
		distributor.LastName = req.LastName
	}
	if req.Phone != "" {
		distributor.Phone = req.Phone
	}
	if req.Address != "" {
		distributor.Address = req.Address
	}
	if req.City != "" {
		distributor.City = req.City
	}
	if req.State != "" {
		distributor.State = req.State
	}
	if req.Country != "" {
		distributor.Country = req.Country
	}
	if req.ZipCode != "" {
		distributor.ZipCode = req.ZipCode
	}
	
	if err := ctrl.distributorService.Update(distributor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, distributor)
}

// List godoc
// @Summary List distributors
// @Tags distributor
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/distributors [get]
func (ctrl *DistributorController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	offset := (page - 1) * limit
	
	distributors, total, err := ctrl.distributorService.List(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data":  distributors,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// GetDownlines godoc
// @Summary Get distributor downlines
// @Tags distributor
// @Produce json
// @Security BearerAuth
// @Param id path int true "Distributor ID"
// @Success 200 {object} []domain.Distributor
// @Router /api/v1/distributors/{id}/downlines [get]
func (ctrl *DistributorController) GetDownlines(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	downlines, err := ctrl.distributorService.GetDownlines(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, downlines)
}

// GetTreeStructure godoc
// @Summary Get tree structure
// @Tags distributor
// @Produce json
// @Security BearerAuth
// @Param id path int true "Distributor ID"
// @Param depth query int false "Tree depth" default(3)
// @Success 200 {object} domain.TreeNode
// @Router /api/v1/distributors/{id}/tree [get]
func (ctrl *DistributorController) GetTreeStructure(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	depth, _ := strconv.Atoi(c.DefaultQuery("depth", "3"))
	
	tree, err := ctrl.distributorService.GetTreeStructure(uint(id), depth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, tree)
}

// AddMemberToTree godoc
// @Summary Add member to tree
// @Tags distributor
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param member body AddMemberRequest true "Member data"
// @Success 201 {object} domain.Distributor
// @Router /api/v1/distributors/add-member [post]
func (ctrl *DistributorController) AddMemberToTree(c *gin.Context) {
	var req AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	member := &domain.Distributor{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		PackageID: req.PackageID,
	}
	
	if err := ctrl.distributorService.AddMemberToTree(member, req.SponsorID, req.Position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, member)
}

// Request/Response DTOs
type RegisterRequest struct {
	FirstName string            `json:"first_name" binding:"required"`
	LastName  string            `json:"last_name" binding:"required"`
	Email     string            `json:"email" binding:"required,email"`
	Password  string            `json:"password" binding:"required,min=6"`
	Phone     string            `json:"phone"`
	Address   string            `json:"address"`
	City      string            `json:"city"`
	State     string            `json:"state"`
	Country   string            `json:"country"`
	ZipCode   string            `json:"zip_code"`
	SponsorID *uint             `json:"sponsor_id"`
	TreeType  domain.TreeType   `json:"tree_type"`
	Position  string            `json:"position"`
	PackageID *uint             `json:"package_id"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	ZipCode   string `json:"zip_code"`
}

type AddMemberRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	SponsorID uint   `json:"sponsor_id" binding:"required"`
	Position  string `json:"position" binding:"required"`
	PackageID *uint  `json:"package_id"`
}
