package handlers

import (
	"net/http"

	"github.com/aswinap13/models"
	"github.com/aswinap13/services"
	"github.com/gin-gonic/gin"
)

type ArrayHandler struct {
	service *services.ArrayService
}

func NewArrayHandler(service *services.ArrayService) *ArrayHandler {
	return &ArrayHandler{service: service}
}

func (h *ArrayHandler) FindPairsWithSum(ctx *gin.Context) {
	var request models.ArrayModel
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.FindPairsWithSum(request.Numbers, request.Target)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"solutions": result})
}
