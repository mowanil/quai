package service

import (
	"kwai/entity"
	repository "kwai/repository/user"
)

type (
	listUserService struct{}
	ListUserService interface {
		ListUsers() ([]entity.User, error)
		FindByName(username string) (entity.User, error)
	}
)

var listUserRepo repository.ListUserRepository

func NewListUserService(repository repository.ListUserRepository) ListUserService {
	listUserRepo = repository
	return &listUserService{}
}

func (*listUserService) ListUsers() ([]entity.User, error) {
	listUserRepo.ListUsers()
	return nil, nil
}

func (*listUserService) FindByName(username string) (entity.User, error) {
	return listUserRepo.FindByName(username)
}
