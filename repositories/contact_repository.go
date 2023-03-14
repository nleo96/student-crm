package repositories

import (
	"open-crm-api/errors"
	"open-crm-api/models"

	"gorm.io/gorm"
)

type ContactRepository struct {
	database *gorm.DB
}

func NewContactRepository(database *gorm.DB) *ContactRepository {
	return &ContactRepository{database: database}
}

func (u ContactRepository) FindByID(
	id uint,
	mode ScanMode,
) (*models.Contact, error) {
	if id == 0 {
		return nil, errors.ErrDataNotFound
	}

	var contact *models.Contact = &models.Contact{}
	var result *gorm.DB

	switch mode {
	case ScanModeSoft:
		result = u.database.Where("id = ?", id).First(contact)
	case ScanModeDeep:
		result = u.database.Preload("Company").Where("id = ?", id).First(contact)
	default:
		return nil, errors.ErrInvalidOption
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.ErrDataNotFound.Wrap(result.Error)
		}

		return nil, errors.ErrDatabase.Wrap(result.Error)
	}

	return contact, nil
}

func (u ContactRepository) Create(
	contact *models.Contact,
) (*models.Contact, error) {
	result := u.database.Create(contact)

	if result.Error != nil {
		return nil, errors.ErrDatabase.Wrap(result.Error)
	}

	result.Scan(contact)

	return contact, nil
}
