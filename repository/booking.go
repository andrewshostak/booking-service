package repository

import (
	"github.com/andrewshostak/booking-service/service"
	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create() (interface{}, error) {
	return nil, nil
}

func (r *bookingRepository) List() ([]service.Booking, error) {
	var dbBookings []Booking
	if result := r.db.Find(&dbBookings); result.Error != nil {
		return nil, result.Error
	}

	return toServiceBookings(dbBookings), nil
}

func (r *bookingRepository) Delete() (interface{}, error) {
	return nil, nil
}
