package service

import (
	"kwai/entity"
	repository "kwai/repository/user"
)

type (
	registrationService struct{}
	RegistrationService interface {
		RegisterUser(user *entity.User) (resultUser *entity.User, err error)
		UnregisterUser(user *entity.User) error
		ModifyUser(user *entity.User) error
	}
)

var registrationRepo repository.RegistrationRepository

func NewRegistrationService(repository repository.RegistrationRepository) RegistrationService {
	registrationRepo = repository
	return &registrationService{}
}

func (*registrationService) RegisterUser(user *entity.User) (*entity.User, error) {
	u, err := registrationRepo.RegisterUser(user)
	return u, err
}

func (*registrationService) UnregisterUser(user *entity.User) error {
	err := registrationRepo.UnregisterUser(user)
	return err
}

func (*registrationService) ModifyUser(user *entity.User) error {
	err := registrationRepo.ModifyUser(user)
	return err
}
