package services

import (
	"open-crm-api/errors"
	"open-crm-api/models"
	"open-crm-api/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repository *repositories.UserRepository
}

func (u UserService) Authenticate(
	credentials *models.UserCredentials,
) (*models.User, error) {
	var user *models.User
	var err error

	user, err = u.repository.FindByEmail(credentials.UserName)

	if err != nil {
		if errors.Is(errors.Unwrap(err), gorm.ErrRecordNotFound) {
			return nil, errors.ErrInvalidCredentials.Wrap(err)
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	)

	if err != nil {
		return nil, errors.ErrInvalidCredentials.Wrap(err)
	}

	return user, nil
}

func (u UserService) SignUp(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		14,
	)

	if err != nil {
		return nil, errors.ErrHashing.Wrap(err)
	}

	user.Password = string(hashedPassword)
	return u.repository.Create(user)

}

func NewUserService(repository *repositories.UserRepository) *UserService {
	service := &UserService{repository: repository}
	return service
}
