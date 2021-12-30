package addresses

import (
	"CalFit/business/addresses"
	"context"

	"gorm.io/gorm"
)

type AddressRepository struct {
	Conn *gorm.DB
}

func NewAddressRepository(conn *gorm.DB) addresses.DomainRepository {
	return &AddressRepository{Conn: conn}
}

func (b *AddressRepository) GetAll(ctx context.Context) ([]addresses.Domain, error) {
	var addressesModel []Address
	if err := b.Conn.Find(&addressesModel).Error; err != nil {
		return nil, err
	}
	var result []addresses.Domain = ToListDomain(addressesModel)
	return result, nil
}

func (b *AddressRepository) GetById(ctx context.Context, id string) (addresses.Domain, error) {
	var address Address
	if err := b.Conn.Where("id = ?", id).First(&address).Error; err != nil {
		return addresses.Domain{}, err
	}
	return address.ToDomain(), nil
}

func (b *AddressRepository) Create(ctx context.Context, address addresses.Domain) (addresses.Domain, error) {
	createdAddress := Address{
		Address:   address.Address,
		District: address.District,
		City:      address.City,
		Postal_code: address.Postal_code,
	}
	
	insertErr := b.Conn.Create(&createdAddress).Error
	if insertErr != nil {
		return addresses.Domain{}, insertErr
	}
	return createdAddress.ToDomain(), nil
}
