package service

import (
	"github.com/andrewshostak/booking-service/handler"
	"time"
)

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
		Birthday:      b.Birthday,
		LaunchpadId:   b.LaunchpadId,
		DestinationId: b.DestinationId,
		LaunchDate:    b.LaunchDate,
	}
}
