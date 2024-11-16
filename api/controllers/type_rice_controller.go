package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/enums"
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
	role, ok := GetRole(c)
	if !ok {
		return
	}
	if role != enums.ROLE_ADMIN {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access is restricted to administrators only"})
		return
	}
	err := u.typeRice.AddTypeRice(c, userID, &req)
	if err != nil {
		RespondError(c, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(c, nil)
}
