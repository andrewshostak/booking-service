package repository

import (
	"github.com/andrewshostak/booking-service/service"
	"time"
)

type BookingRepository interface {
	Create(toCreate service.BookingToCreate) (*service.Booking, error)
	List() ([]service.Booking, error)
	Delete(id uint) error
}

type LaunchpadRepository interface {
	IsLaunchpadAvailable(launchpadId string, launchDate time.Time) (bool, error)
}
