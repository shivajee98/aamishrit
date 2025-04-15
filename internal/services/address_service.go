package services

import (
	"github.com/shivajee98/aamishrit/internal/model"
	"github.com/shivajee98/aamishrit/internal/repository"
)

type AddressService interface {
	CreateAddress(address *model.Address) error
	GetAddressesByUserID(clerkID string) ([]model.Address, error)
	GetAddressByID(addressID uint) (*model.Address, error)
	UpdateAddress(address *model.Address) error
	DeleteAddress(addressID uint) error
	SetDefaultAddress(userID uint, addressID uint) error
	GetDefaultAddress(userID uint) (*model.Address, error)
}

type addressService struct {
	repo repository.AddressRepository
}

func InitAddressService(repo repository.AddressRepository) AddressService {
	return &addressService{repo: repo}
}

func (s *addressService) CreateAddress(address *model.Address) error {
	return s.repo.CreateAddress(address)
}

func (s *addressService) GetAddressesByUserID(clerkID string) ([]model.Address, error) {
	return s.repo.GetAddressesByUserID(clerkID)
}

func (s *addressService) GetAddressByID(addressID uint) (*model.Address, error) {
	return s.repo.GetAddressByID(addressID)
}

func (s *addressService) UpdateAddress(address *model.Address) error {
	return s.repo.UpdateAddress(address)
}

func (s *addressService) DeleteAddress(addressID uint) error {
	return s.repo.DeleteAddress(addressID)
}

func (s *addressService) SetDefaultAddress(userID uint, addressID uint) error {
	return s.repo.SetDefaultAddress(userID, addressID)
}

func (s *addressService) GetDefaultAddress(userID uint) (*model.Address, error) {
	return s.repo.GetDefaultAddress(userID)
}
