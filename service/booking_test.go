package service_test

import (
	"errors"
	"github.com/andrewshostak/booking-service/handler"
	"github.com/andrewshostak/booking-service/service"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBookingService_Create(t *testing.T) {
	bookingMock := &service.BookingRepositoryMock{}
	launchpadMock := &service.LaunchpadRepositoryMock{}
	bookingService := service.NewBookingService(bookingMock, launchpadMock)

	toCreate := handler.BookingToCreate{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Gender:        gofakeit.Gender(),
		Birthday:      gofakeit.Date().Format(service.DateFormat),
		LaunchpadId:   gofakeit.UUID(),
		DestinationId: gofakeit.UUID(),
		LaunchDate:    gofakeit.Date().Format(service.DateFormat),
	}
	errFromRepo := errors.New(gofakeit.Sentence(3))
	expectedDateParam, _ := time.Parse(service.DateFormat, toCreate.LaunchDate)
	expectedBirthdayParam, _ := time.Parse(service.DateFormat, toCreate.Birthday)
	expectedCreateParam := &service.BookingToCreate{
		FirstName:     toCreate.FirstName,
		LastName:      toCreate.LastName,
		Gender:        toCreate.Gender,
		Birthday:      expectedBirthdayParam,
		LaunchpadId:   toCreate.LaunchpadId,
		DestinationId: toCreate.DestinationId,
		LaunchDate:    expectedDateParam,
	}

	t.Run("error when parsing date", func(t *testing.T) {
		toCreateCopy := toCreate
		toCreateCopy.Birthday = "wrong"

		result, err := bookingService.Create(toCreateCopy)
		assert.Nil(t, result)
		assert.Error(t, err)

		anotherCopy := toCreate
		anotherCopy.LaunchDate = "wrong"

		result, err = bookingService.Create(anotherCopy)
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("error when getting launchpad availability", func(t *testing.T) {
		launchpadMock.On("IsLaunchpadAvailable", toCreate.LaunchpadId, expectedDateParam).Return(false, errFromRepo).Once()

		result, err := bookingService.Create(toCreate)
		assert.Nil(t, result)
		assert.Equal(t, errFromRepo, err)

		launchpadMock.AssertExpectations(t)
	})

	t.Run("error when launchpad is not available", func(t *testing.T) {
		errLaunchpad := errors.New("launchpad is not available")
		launchpadMock.On("IsLaunchpadAvailable", toCreate.LaunchpadId, expectedDateParam).Return(false, nil).Once()

		result, err := bookingService.Create(toCreate)
		assert.Nil(t, result)
		assert.Equal(t, errLaunchpad, err)

		launchpadMock.AssertExpectations(t)
	})

	t.Run("error when creating booking", func(t *testing.T) {
		launchpadMock.On("IsLaunchpadAvailable", toCreate.LaunchpadId, expectedDateParam).Return(true, nil).Once()
		bookingMock.On("Create", *expectedCreateParam).Return(nil, errFromRepo).Once()

		result, err := bookingService.Create(toCreate)
		assert.Nil(t, result)
		assert.Equal(t, errFromRepo, err)

		launchpadMock.AssertExpectations(t)
		bookingMock.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		expectedServiceBooking := service.Booking{
			Id:            uint(gofakeit.Uint8()),
			FirstName:     expectedCreateParam.FirstName,
			LastName:      expectedCreateParam.LastName,
			Gender:        expectedCreateParam.Gender,
			Birthday:      expectedCreateParam.Birthday,
			LaunchpadId:   expectedCreateParam.LaunchpadId,
			DestinationId: expectedCreateParam.DestinationId,
			LaunchDate:    expectedCreateParam.LaunchDate,
		}

		launchpadMock.On("IsLaunchpadAvailable", toCreate.LaunchpadId, expectedDateParam).Return(true, nil).Once()
		bookingMock.On("Create", *expectedCreateParam).Return(&expectedServiceBooking, nil).Once()

		expectedResult := handler.Booking{
			Id:            expectedServiceBooking.Id,
			FirstName:     expectedServiceBooking.FirstName,
			LastName:      expectedServiceBooking.LastName,
			Gender:        expectedServiceBooking.Gender,
			Birthday:      toCreate.Birthday,
			LaunchpadId:   expectedServiceBooking.LaunchpadId,
			DestinationId: expectedServiceBooking.DestinationId,
			LaunchDate:    toCreate.LaunchDate,
		}

		result, err := bookingService.Create(toCreate)
		assert.NoError(t, err)
		assert.Equal(t, &expectedResult, result)

		launchpadMock.AssertExpectations(t)
		bookingMock.AssertExpectations(t)
	})
}

func TestBookingService_List(t *testing.T) {
	bookingMock := &service.BookingRepositoryMock{}
	launchpadMock := &service.LaunchpadRepositoryMock{}
	bookingService := service.NewBookingService(bookingMock, launchpadMock)

	t.Run("error when getting bookings", func(t *testing.T) {
		errFromRepo := errors.New(gofakeit.Sentence(3))
		bookingMock.On("List").Return(nil, errFromRepo).Once()

		result, err := bookingService.List()
		assert.Nil(t, result)
		assert.Equal(t, errFromRepo, err)

		bookingMock.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		expectedServiceBookings := []service.Booking{
			{
				Id:            uint(gofakeit.Uint8()),
				FirstName:     gofakeit.FirstName(),
				LastName:      gofakeit.LastName(),
				Gender:        gofakeit.Gender(),
				Birthday:      gofakeit.Date(),
				LaunchpadId:   gofakeit.UUID(),
				DestinationId: gofakeit.UUID(),
				LaunchDate:    gofakeit.Date(),
			},
		}

		expectedResult := []handler.Booking{
			{
				Id:            expectedServiceBookings[0].Id,
				FirstName:     expectedServiceBookings[0].FirstName,
				LastName:      expectedServiceBookings[0].LastName,
				Gender:        expectedServiceBookings[0].Gender,
				Birthday:      expectedServiceBookings[0].Birthday.Format(service.DateFormat),
				LaunchpadId:   expectedServiceBookings[0].LaunchpadId,
				DestinationId: expectedServiceBookings[0].DestinationId,
				LaunchDate:    expectedServiceBookings[0].LaunchDate.Format(service.DateFormat),
			},
		}

		bookingMock.On("List").Return(expectedServiceBookings, nil).Once()

		result, err := bookingService.List()
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)

		bookingMock.AssertExpectations(t)
	})
}

func TestBookingService_Delete(t *testing.T) {
	bookingMock := &service.BookingRepositoryMock{}
	launchpadMock := &service.LaunchpadRepositoryMock{}
	bookingService := service.NewBookingService(bookingMock, launchpadMock)
	id := uint(gofakeit.Uint8())

	t.Run("error when deleting booking", func(t *testing.T) {
		errFromRepo := errors.New(gofakeit.Sentence(3))
		bookingMock.On("Delete", id).Return(errFromRepo).Once()

		err := bookingService.Delete(id)
		assert.Equal(t, errFromRepo, err)

		bookingMock.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		bookingMock.On("Delete", id).Return(nil).Once()

		err := bookingService.Delete(id)
		assert.NoError(t, err)

		bookingMock.AssertExpectations(t)
	})
}
