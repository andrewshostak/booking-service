package service

type bookingService struct {
}

type BookingService interface {
	Create() (interface{}, error)
	List() (interface{}, error)
	Delete() (interface{}, error)
}

func NewBookingService() BookingService {
	return &bookingService{}
}

func (s *bookingService) Create() (interface{}, error) {
	return nil, nil
}

func (s *bookingService) List() (interface{}, error) {
	return nil, nil
}

func (s *bookingService) Delete() (interface{}, error) {
	return nil, nil
}
