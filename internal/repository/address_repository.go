package repository

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address *model.Address) error
	GetAddressesByUserID(clerkID string) ([]model.Address, error)
	GetAddressByID(addressID uint) (*model.Address, error)
	UpdateAddress(address *model.Address) error
	DeleteAddress(addressID uint) error
	SetDefaultAddress(userID uint, addressID uint) error
	GetDefaultAddress(userID uint) (*model.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func InitAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

func (r *addressRepository) CreateAddress(address *model.Address) error {
	return r.db.Create(address).Error
}

// func (r *addressRepository) GetAddressesByUserID(clerkID string) ([]model.Address, error) {
// 	var addresses []model.Address
// 	err := r.db.Where("user_id = ?", clerkID).Find(&addresses).Error
// 	return addresses, err
// }

func (r *addressRepository) GetAddressesByUserID(clerkID string) ([]model.Address, error) {
	var user model.User
	if err := r.db.Where("clerk_id = ?", clerkID).First(&user).Error; err != nil {
		return nil, err
	}

	var addresses []model.Address
	err := r.db.Where("user_id = ?", user.ID).Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) GetAddressByID(addressID uint) (*model.Address, error) {
	var address model.Address
	err := r.db.First(&address, addressID).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *addressRepository) UpdateAddress(address *model.Address) error {
	return r.db.Model(&model.Address{}).Where("id = ?", address.ID).Updates(address).Error
}

func (r *addressRepository) DeleteAddress(addressID uint) error {
	return r.db.Delete(&model.Address{}, addressID).Error
}

func (r *addressRepository) SetDefaultAddress(userID uint, addressID uint) error {
	// First unset previous default address (if any)
	err := r.db.Model(&model.Address{}).
		Where("user_id = ?", userID).
		Update("is_default", false).Error
	if err != nil {
		return err
	}

	// Then set new default address
	return r.db.Model(&model.Address{}).
		Where("id = ? AND user_id = ?", addressID, userID).
		Update("is_default", true).Error
}

func (r *addressRepository) GetDefaultAddress(userID uint) (*model.Address, error) {
	var address model.Address
	err := r.db.Where("user_id = ? AND is_default = ?", userID, true).
		First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}
