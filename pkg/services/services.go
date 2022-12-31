package services

import "fatawa-api/pkg/models"

type FatwaService interface {
	CreateFatwa(fatwa *models.Fatwa) (*models.FatwaDb, error)
	UpdateFatwa(string, *models.FatwaDb) (*models.FatwaDb, error)
	FindFatwaById(string) (*models.FatwaDb, error)
	FindFatawa(page int, limit int) ([]*models.FatwaDb, error)
	DeleteFatwa(string) error
}
