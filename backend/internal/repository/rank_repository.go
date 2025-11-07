package repository

import (
	"errors"

	"github.com/mlm-app/backend/internal/domain"
	"gorm.io/gorm"
)

type RankRepository interface {
	Create(rank *domain.Rank) error
	FindByID(id uint) (*domain.Rank, error)
	FindByName(name string) (*domain.Rank, error)
	Update(rank *domain.Rank) error
	Delete(id uint) error
	List() ([]domain.Rank, error)
}

type rankRepository struct {
	db *gorm.DB
}

func NewRankRepository(db *gorm.DB) RankRepository {
	return &rankRepository{db: db}
}

func (r *rankRepository) Create(rank *domain.Rank) error {
	return r.db.Create(rank).Error
}

func (r *rankRepository) FindByID(id uint) (*domain.Rank, error) {
	var rank domain.Rank
	err := r.db.First(&rank, id).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rank not found")
		}
		return nil, err
	}
	return &rank, nil
}

func (r *rankRepository) FindByName(name string) (*domain.Rank, error) {
	var rank domain.Rank
	err := r.db.Where("name = ?", name).First(&rank).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rank not found")
		}
		return nil, err
	}
	return &rank, nil
}

func (r *rankRepository) Update(rank *domain.Rank) error {
	return r.db.Save(rank).Error
}

func (r *rankRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Rank{}, id).Error
}

func (r *rankRepository) List() ([]domain.Rank, error) {
	var ranks []domain.Rank
	err := r.db.Order("level ASC").Find(&ranks).Error
	return ranks, err
}
