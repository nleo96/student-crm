package repositories

import "open-crm-api/models"

type ScanMode uint

const (
	ScanModeSoft ScanMode = (iota + 1) << 32
	ScanModeDeep
)

type Repository interface {
	FindByID(id int) (*models.Model, error)
	Find(filter any) ([]*models.Model, error)
}
