package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"
)

type TypeRiceController struct {
	typeRice *services.TypeRiceService
}

func NewTypeRiceController(typeRice *services.TypeRiceService) *TypeRiceController {
	return &TypeRiceController{
		typeRice: typeRice,
	}
}

func (u *TypeRiceController) AddTypeRice(c *gin.Context) {
	var req entities.CreateTypeRiceRequest
	if !BindAndValidate(c, &req) {
		return
	}
	userID, ok := GetUserID(c)
	if !ok {
		return
	}
	err := u.typeRice.AddTypeRice(c, userID, &req)
	if err != nil {
		RespondError(c, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(c, nil)
}
