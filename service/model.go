package service

import (
	"github.com/andrewshostak/booking-service/handler"
	"time"
)

type BookingToCreate struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadId   string
	DestinationId string
	LaunchDate    time.Time
}

func toBookingCreation(bookingToCreate handler.BookingToCreate) (*BookingToCreate, error) {
	birthday, err := time.Parse(dateFormat, bookingToCreate.Birthday)
	if err != nil {
		return nil, err
	}

	return &BookingToCreate{
		FirstName:     bookingToCreate.FirstName,
		LastName:      bookingToCreate.LastName,
		Gender:        bookingToCreate.Gender,
		Birthday:      birthday,
		LaunchpadId:   bookingToCreate.LaunchpadId,
		DestinationId: bookingToCreate.DestinationId,
		LaunchDate:    bookingToCreate.LaunchDate,
	}, nil
}

type Booking struct {
	Id            uint
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadId   string
	DestinationId string
	LaunchDate    time.Time
}

func (b Booking) toHandlerModel() handler.Booking {
	return handler.Booking{
		Id:            b.Id,
		FirstName:     b.FirstName,
		LastName:      b.LastName,
		Gender:        b.Gender,
		Birthday:      b.Birthday.Format(dateFormat),
		LaunchpadId:   b.LaunchpadId,
		DestinationId: b.DestinationId,
		LaunchDate:    b.LaunchDate,
	}
}

func toHandlerBookings(bookings []Booking) []handler.Booking {
	handlerBookings := make([]handler.Booking, 0, len(bookings))
	for i := range bookings {
		handlerBookings = append(handlerBookings, bookings[i].toHandlerModel())
	}
	return handlerBookings
}

const dateFormat = "2006-01-02"
