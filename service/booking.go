package service

import (
	"errors"
	"github.com/andrewshostak/booking-service/handler"
)

type bookingService struct {
	br ListerDeleterCreator
	lr LaunchpadChecker
}

func NewBookingService(br ListerDeleterCreator, lr LaunchpadChecker) BookingService {
	return &bookingService{br: br, lr: lr}
}

func (s *bookingService) Create(toCreate handler.BookingToCreate) (*handler.Booking, error) {
	booking, err := toBookingCreation(toCreate)
	if err != nil {
		return nil, err
	}

	isAvailable, err := s.lr.IsLaunchpadAvailable(booking.LaunchpadId, booking.LaunchDate)
	if err != nil {
		return nil, err
	}

	if !isAvailable {
		return nil, errors.New("launchpad is not available")
	}

	created, err := s.br.Create(*booking)
	if err != nil {
		return nil, err
	}

	handlerModel := created.toHandlerModel()
	return &handlerModel, nil
}

func (s *bookingService) List() ([]handler.Booking, error) {
	list, err := s.br.List()
	if err != nil {
		return nil, err
	}

	return toHandlerBookings(list), nil
}

func (s *bookingService) Delete(id uint) error {
	return s.br.Delete(id)
}
