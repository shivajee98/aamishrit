package services

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type UserService interface {
	RegisterUser(user *model.User) error
	GetUserByPhone(phone string) (*model.User, error)
	UpdateUser(user *model.User) error
	GetUserByClerkID(clerkID string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func InitUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetUser implements UserService.
func (u *userService) GetUserByPhone(phone string) (*model.User, error) {
	return u.userRepo.GetUserByPhone(phone)
}

// RegisterUser implements UserService.
func (u *userService) RegisterUser(user *model.User) error {
	return u.userRepo.RegisterUser(user)
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(user *model.User) error {
	return u.userRepo.UpdateUser(user)
}

// Get User By Clerk Id
func (u *userService) GetUserByClerkID(clerkID string) (*model.User, error) {
	return u.userRepo.GetUserByClerkID(clerkID)
}