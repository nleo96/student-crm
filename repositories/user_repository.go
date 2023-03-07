package repositories

import (
	"fmt"
	"open-crm-api/failures"
	"open-crm-api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func (u UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := u.database.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, failures.ErrDataNotFound.Wrap(result.Error)
		} else {
			return nil, failures.ErrDatabaseOperation.Wrap(result.Error)
		}
	}

	return &user, nil
}

func (u UserRepository) Create(user *models.User) (*models.User, error) {
	result := u.database.Create(user)

	if result.Error != nil {
		return nil, fmt.Errorf(
			"%w: %v",
			failures.ErrPasswordHashGeneration,
			result.Error,
		)

	}

	result.Scan(user)

	return user, nil
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database: database}
}
