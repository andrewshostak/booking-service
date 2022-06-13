package service

import "github.com/andrewshostak/booking-service/handler"

type bookingService struct {
	br ListerDeleterCreator
}

func NewBookingService(br ListerDeleterCreator) BookingService {
	return &bookingService{br: br}
}

func (s *bookingService) Create(toCreate handler.BookingToCreate) (*handler.Booking, error) {
	booking, err := toBookingCreation(toCreate)
	if err != nil {
		return nil, err
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
