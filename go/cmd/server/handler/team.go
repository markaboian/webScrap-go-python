package handler

import (
	"go-python/internal/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service product.IService
}

func (h *Handler) GetTeamByRank(ctx *gin.Context) {
	rankParam := ctx.Param("rank")
	rank, err := strconv.Atoi(rankParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while converting rank - " + err.Error(),
		})
		return
	}

	team, err := h.Service.GetTeamByRank(rank);
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Error: team not found - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, team)
}