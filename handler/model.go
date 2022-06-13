package handler

import "time"

type Booking struct {
	Id            uint      `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        string    `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	LaunchpadId   string    `json:"launchpad_id"`
	DestinationId string    `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
}
