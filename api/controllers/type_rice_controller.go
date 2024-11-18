package controllers

import (
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"

	"github.com/gin-gonic/gin"
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

	code, err := u.typeRice.AddTypeRice(c, userID, &req)
	if err != nil {
		RespondError(c, http.StatusInternalServerError, err)
		return
	}
	if code != nil {
		RespondConflict(c, http.StatusOK, "type rice name already exists")
		return
	}
	RespondSuccess(c, nil)
}

func (u *TypeRiceController) GetTypeRice(c *gin.Context) {
	userID, ok := GetUserID(c)
	if !ok {
		return
	}
	resp, err := u.typeRice.GetTypeRiceNameByUserID(c, userID)
	if err != nil {
		RespondError(c, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(c, resp)
}

func (u *TypeRiceController) DeleteById(ctx *gin.Context) {
	id := GetIdFromParam(ctx, "id")
	err := u.typeRice.DeleteById(ctx, id)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, nil)
}
