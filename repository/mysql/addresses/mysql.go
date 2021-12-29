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
	createdAddress.BeforeCreate()
	
	insertErr := b.Conn.Create(&createdAddress).Error
	if insertErr != nil {
		return addresses.Domain{}, insertErr
	}
	return createdAddress.ToDomain(), nil
}

// func (b *AddressRepository) UpdateStatus(ctx context.Context, id string, status bool) (addresses.Domain, error) {
// 	var book addresses
// 	if err := b.Conn.Where("book_id = ?", id).First(&book).Error; err != nil {
// 		return addresses.Domain{}, err
// 	}
// 	book.Status = status
// 	book.UpdatedAt = time.Now()
// 	if err := b.Conn.Save(&book).Error; err != nil {
// 		return addresses.Domain{}, err
// 	}
// 	return book.ToDomain(), nil
// }

// // func (b *AddressRepository) Update(user *User) error {
// // 	return b.Conn.Save(user).Error
// // }

// // func (b *AddressRepository) Delete(user *User) error {
// // 	return b.Conn.Delete(user).Error
// // }
