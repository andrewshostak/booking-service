package repository

import "gorm.io/gorm"

type bookingRepository struct {
	db *gorm.DB
}

type BookingRepository interface {
	Create() (interface{}, error)
	List() (interface{}, error)
	Delete() (interface{}, error)
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create() (interface{}, error) {
	return nil, nil
}

func (r *bookingRepository) List() (interface{}, error) {
	return nil, nil
}

func (r *bookingRepository) Delete() (interface{}, error) {
	return nil, nil
}
