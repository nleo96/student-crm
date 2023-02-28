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
			return nil, fmt.Errorf(
				"%w: %v",
				failures.ErrNotFound,
				result.Error,
			)
		}

		return nil, result.Error
	}

	return &user, nil
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database: database}
}
