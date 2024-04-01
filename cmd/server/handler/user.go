package handler

import (
	"UsersAPI/internal/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service product.IService
}

func (h *Handler) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"Error al convertir el id - " + err.Error(),
		})
		return
	}

	user, err := h.Service.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message:": "Error: usuario no encontrado - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}