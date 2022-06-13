package service

import "github.com/andrewshostak/booking-service/handler"

type BookingService interface {
	Create() (interface{}, error)
	List() ([]handler.Booking, error)
	Delete() (interface{}, error)
}

type Lister interface {
	List() ([]Booking, error)
}
