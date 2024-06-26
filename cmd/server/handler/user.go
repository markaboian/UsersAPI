package handler

import (
	"UsersAPI/internal/domain"
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
			"message": "Error: while converting id - " + err.Error(),
		})
		return
	}

	user, err := h.Service.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message:": "Error: user not found - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) AddUser(ctx *gin.Context) {
	var newUser domain.User

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding JSON - " + err.Error(),
		})
		return
	}

	user, err := h.Service.AddUser(newUser.Name, newUser.LastName, newUser.PlaceOfBirth, newUser.DateOfBirth, newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while adding user - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var userUpdated domain.User

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while converting id - " + err.Error(),
		})
		return
	}

	if err := ctx.BindJSON(&userUpdated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding JSON - " + err.Error(),
		})
		return
	}

	user, err := h.Service.UpdateUser(id, &userUpdated)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while updating user - " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, user)

}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while converting id - " + err.Error(),
		})
		return
	}

	if err := h.Service.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while deleting user - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}

// Products methods

func (h *Handler) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while converting id - " + err.Error(),
		})
		return
	}

	product, err := h.Service.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Error: product not found - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Handler) AddProduct(ctx *gin.Context) {
	var newProduct domain.Product

	if err := ctx.BindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while binding json - " + err.Error(), 
		})
		return
	}

	product, err := h.Service.AddProduct(newProduct.Name, newProduct.Price, newProduct.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while adding product - " + err.Error(), 
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (h *Handler) GetProductsByUserId(ctx *gin.Context) {
	idParam := ctx.Param("id")
	idUser, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while converting id - " + err.Error(),
		})
		return
	}

	productsWithUser, err := h.Service.GetProductsByUserId(idUser)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Error: products not found - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, productsWithUser)
}

func (h *Handler) UpdateProduct(ctx *gin.Context) {
	var productUpdated domain.Product

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while converting id - " + err.Error(), 
		})
		return
	}

	if err := ctx.BindJSON(&productUpdated); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while binding json - " + err.Error(),
		})
		return
	}

	product, err := h.Service.UpdateProduct(id, &productUpdated)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while updating product - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (h *Handler) DeleteProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message" : "Error: while converting id - " + err.Error(),
		})
		return
	}

	if err := h.Service.DeleteProduct(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error: while deleting the product - " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})


}