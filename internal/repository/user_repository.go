package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *model.User) error
	GetUserByPhone(phone string) (*model.User, error)
	GetUserByClerkID(clerkID string) (*model.User, error)
	UpdateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetUserByClerkID(clerkID string) (*model.User, error) {
	var user model.User
	err := r.db.Where("clerk_id = ?", clerkID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
