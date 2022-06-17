package usecase

import "kwai/entity"

// Provide several ways to list users
type ListUserUseCase interface {
	// List all registered users
	ListUsers() ([]entity.User, error)
	// Find a specific user by his name
	FindByName(username string) (entity.User, error)
}
