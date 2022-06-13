package handler

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type BookingToCreate struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
	Birthday      string `json:"birthday" binding:"required,date"`
	LaunchpadId   string `json:"launchpad_id" binding:"required"`
	DestinationId string `json:"destination_id" binding:"required"`
	LaunchDate    string `json:"launch_date" binding:"required,date"`
}

type Booking struct {
	Id            uint   `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	LaunchpadId   string `json:"launchpad_id"`
	DestinationId string `json:"destination_id"`
	LaunchDate    string `json:"launch_date"`
}

var ValidateDate validator.Func = func(fl validator.FieldLevel) bool {
	search, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if len(search) == 0 {
		return false
	}

	matched, err := regexp.MatchString("\\d{4}-\\d{2}-\\d{2}", search)
	if err != nil {
		return false
	}

	return matched
}
