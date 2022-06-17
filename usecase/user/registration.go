package usecase

import "kwai/entity"

type RegistrationUseCase interface {
	RegisterUser(user *entity.User) (resultUser *entity.User, err error) // same as creating an user account
	UnregisterUser(user *entity.User) error                              // same as deleting an user account
	ModifyUser(user *entity.User) error                                  // same as editing/updating user account details
}
