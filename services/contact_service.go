package services

import (
	"open-crm-api/repositories"
)

type ContactService struct {
	repository *repositories.ContactRepository
}

func (s ContactService) Repository() *repositories.ContactRepository {
	return s.repository
}

func NewContactService(
	repository *repositories.ContactRepository,
) *ContactService {
	return &ContactService{repository: repository}
}
