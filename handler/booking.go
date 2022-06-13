package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type bookingHandler struct {
	bs ListerDeleter
}

func NewBookingHandler(bs ListerDeleter) BookingHandler {
	return &bookingHandler{bs: bs}
}

func (h *bookingHandler) Create(context *gin.Context) {

}

func (h *bookingHandler) List(context *gin.Context) {
	list, err := h.bs.List()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"bookings": list})
}

func (h *bookingHandler) Delete(context *gin.Context) {
	var uriParams struct {
		Id uint `uri:"id" binding:"required"`
	}
	if err := context.ShouldBindUri(&uriParams); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.bs.Delete(uriParams.Id); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
