package repository

import (
	"log"

	"kwai/entity"
	usecase "kwai/usecase/user"
)

type listUserRepository struct{}

type ListUserRepository interface {
	usecase.ListUserUseCase
}

func NewListUserRepository() ListUserRepository {
	return &listUserRepository{}
}

func (*listUserRepository) ListUsers() ([]entity.User, error) {
	log.Println("Retrieving all users from database")
	return nil, nil
}

func (*listUserRepository) FindByName(username string) (entity.User, error) {
	var user entity.User
	log.Printf("Retrieving user with name %s from database\n", username)
	return user, nil
}
