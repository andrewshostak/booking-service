package repository

import "github.com/andrewshostak/booking-service/service"

type BookingRepository interface {
	Create() (interface{}, error)
	List() ([]service.Booking, error)
	Delete(id uint) error
}
