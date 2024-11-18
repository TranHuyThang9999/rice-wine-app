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
	userID := GetUserID(c)

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
	userID := GetUserID(c)

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

func (u *TypeRiceController) UpdateById(ctx *gin.Context) {
	var req entities.UpdateTypeRiceRequest
	if !BindAndValidate(ctx, &req) {
		return
	}
	code, err := u.typeRice.UpdateById(ctx, GetUserID(ctx), &req)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	if code != nil {
		RespondConflict(ctx, http.StatusOK, "type rice name already exists")
		return
	}
	RespondSuccess(ctx, nil)
}
