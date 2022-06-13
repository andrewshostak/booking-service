package service

type bookingService struct {
	br interface{}
}

type BookingService interface {
	Create() (interface{}, error)
	List() (interface{}, error)
	Delete() (interface{}, error)
}

func NewBookingService(br interface{}) BookingService {
	return &bookingService{br: br}
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
