package repository

import (
	"github.com/andrewshostak/booking-service/service"
	"time"
)

type Booking struct {
	Id            uint      `gorm:"id"`
	FirstName     string    `gorm:"first_name"`
	LastName      string    `gorm:"last_name"`
	Gender        string    `gorm:"gender"`
	Birthday      time.Time `gorm:"birthday"`
	LaunchpadId   string    `gorm:"launchpad_id"`
	DestinationId string    `gorm:"destination_id"`
	LaunchDate    time.Time `gorm:"launch_date"`
}

func (b Booking) toServiceModel() service.Booking {
	return service.Booking{
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

func toServiceBookings(bookings []Booking) []service.Booking {
	serviceBookings := make([]service.Booking, 0, len(bookings))
	for i := range bookings {
		serviceBookings = append(serviceBookings, bookings[i].toServiceModel())
	}
	return serviceBookings
}
