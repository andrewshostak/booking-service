package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type bookingHandler struct {
	bs Lister
}

func NewBookingHandler(bs Lister) BookingHandler {
	return &bookingHandler{bs: bs}
}

func (h *bookingHandler) Create(context *gin.Context) {

}

func (h *bookingHandler) List(context *gin.Context) {
	list, err := h.bs.List()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"bookings": list})
}

func (h *bookingHandler) Delete(context *gin.Context) {

}
