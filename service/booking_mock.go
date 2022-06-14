package service

import (
	"github.com/stretchr/testify/mock"
)

var _ ListerDeleterCreator = &BookingRepositoryMock{}

type BookingRepositoryMock struct {
	mock.Mock
}

func (m *BookingRepositoryMock) Create(toCreate BookingToCreate) (*Booking, error) {
	args := m.Called(toCreate)

	arg := args.Get(0)
	var booking *Booking
	if arg != nil {
		booking = arg.(*Booking)
	}

	return booking, args.Error(1)
}

func (m *BookingRepositoryMock) List() ([]Booking, error) {
	args := m.Called()

	arg := args.Get(0)
	var bookings []Booking
	if arg != nil {
		bookings = arg.([]Booking)
	}

	return bookings, args.Error(1)
}

func (m *BookingRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
