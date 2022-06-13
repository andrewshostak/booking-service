package handler

import "github.com/gin-gonic/gin"

type bookingHandler struct {
}

type BookingHandler interface {
	Create(context *gin.Context)
	List(context *gin.Context)
	Delete(context *gin.Context)
}

func NewBookingHandler() BookingHandler {
	return &bookingHandler{}
}

func (h *bookingHandler) Create(context *gin.Context) {

}

func (h *bookingHandler) List(context *gin.Context) {

}

func (h *bookingHandler) Delete(context *gin.Context) {

}
