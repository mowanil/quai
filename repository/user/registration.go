package repository

import (
	"log"

	"kwai/entity"
	usecase "kwai/usecase/user"
)

type (
	ruc struct{}
)

type RegistrationRepository interface {
	usecase.RegistrationUseCase
}

func NewRegistrationRepository() RegistrationRepository {
	return &ruc{}
}

func (*ruc) RegisterUser(user *entity.User) (*entity.User, error) {
	log.Printf("Trying to register %#v\n", user)
	return nil, nil
}

func (*ruc) UnregisterUser(user *entity.User) error {
	log.Printf("Trying to unregister %#v\n", user)
	return nil
}

func (*ruc) ModifyUser(user *entity.User) error {
	log.Printf("Trying to modify %#v\n", user)
	return nil
}
