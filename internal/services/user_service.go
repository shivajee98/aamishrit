package services

import (
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type UserService interface {
	RegisterUser(user *model.User) error
	LoginUser(phone, password string) (*model.User, error)
	GetUser(phone string) (*model.User, error)
	UpdateUser(user *model.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func InitUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetUser implements UserService.
func (u *userService) GetUser(phone string) (*model.User, error) {
	user.
	panic("unimplemented")
}

// LoginUser implements UserService.
func (u *userService) LoginUser(phone string, password string) (*model.User, error) {
	panic("unimplemented")
}

// RegisterUser implements UserService.
func (u *userService) RegisterUser(user *model.User) error {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(user *model.User) error {
	panic("unimplemented")
}
