package services

import "github.com/shivajee98/aamishrit/internal/model"

type UserService interface {
	RegisterUser(user *model.User) error
	LoginUser(phone, password string) (*model.User, error)

}
