package services

import (
	"errors"
	"fmt"
	"open-crm-api/failures"
	"open-crm-api/models"
	"open-crm-api/repositories"

	"golang.org/x/crypto/bcrypt"
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
		if errors.Is(err, failures.ErrNotFound) {
			return nil, fmt.Errorf(
				"%w: %v",
				failures.ErrInvalidCredentials,
				err,
			)
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password),
	)

	if err != nil {
		return nil, fmt.Errorf(
			"%w: %v",
			failures.ErrInvalidCredentials,
			err,
		)
	}

	return user, nil
}

func (u UserService) SignUp(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		14,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"%w: %v",
			failures.ErrPasswordHashGeneration,
			err,
		)
	}

	user.Password = string(hashedPassword)
	return u.repository.Create(user)

}

func NewUserService(repository *repositories.UserRepository) *UserService {
	service := &UserService{repository: repository}
	return service
}
