package service

import "github.com/andrewshostak/booking-service/handler"

type BookingService interface {
	Create() (interface{}, error)
	List() ([]handler.Booking, error)
	Delete(id uint) error
}

type Lister interface {
	List() ([]Booking, error)
}

type Deleter interface {
	Delete(id uint) error
}

type ListerDeleter interface {
	Lister
	Deleter
}
