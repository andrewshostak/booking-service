package repository

import "github.com/andrewshostak/booking-service/service"

type BookingRepository interface {
	Create(toCreate service.BookingToCreate) (*service.Booking, error)
	List() ([]service.Booking, error)
	Delete(id uint) error
}
