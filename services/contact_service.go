package services

import "open-crm-api/repositories"

type ContactService struct {
	repository *repositories.ContactRepository
}



func NewContactService(
	repository *repositories.ContactRepository,
) *ContactService {
	return &ContactService{repository: repository}
}
