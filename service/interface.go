package service

import "github.com/andrewshostak/booking-service/handler"

type BookingService interface {
	Create(toCreate handler.BookingToCreate) (*handler.Booking, error)
	List() ([]handler.Booking, error)
	Delete(id uint) error
}

type Lister interface {
	List() ([]Booking, error)
}

type Deleter interface {
	Delete(id uint) error
}

type Creator interface {
	Create(toCreate BookingToCreate) (*Booking, error)
}

type ListerDeleterCreator interface {
	Lister
	Deleter
	Creator
}
