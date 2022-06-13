package handler

import "github.com/gin-gonic/gin"

type BookingHandler interface {
	Create(context *gin.Context)
	List(context *gin.Context)
	Delete(context *gin.Context)
}

type Lister interface {
	List() ([]Booking, error)
}

type Deleter interface {
	Delete(id uint) error
}

type Creator interface {
	Create(toCreate BookingToCreate) (*Booking, error)
}

type ListerDeleterCreator interface {
	Lister
	Deleter
	Creator
}
