package service

import "github.com/andrewshostak/booking-service/handler"

type bookingService struct {
	br ListerDeleter
}

func NewBookingService(br ListerDeleter) BookingService {
	return &bookingService{br: br}
}

func (s *bookingService) Create() (interface{}, error) {
	return nil, nil
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
