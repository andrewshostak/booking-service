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
