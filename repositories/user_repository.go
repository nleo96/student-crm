package repositories

import (
	"open-crm-api/errors"
	"open-crm-api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func (u UserRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User = &models.User{}
	result := u.database.Where("email = ?", email).First(user)

	if result.Error != nil {
		return nil, errors.ErrDatabase.Wrap(result.Error)
	}

	return user, nil
}

func (u UserRepository) Create(user *models.User) (*models.User, error) {
	result := u.database.Create(user)

	if result.Error != nil {
		return nil, errors.ErrDatabase.Wrap(result.Error)
	}

	result.Scan(user)

	return user, nil
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database: database}
}
