package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"
)

type RiceController struct {
	rice *services.RiceService
}

func NewRiceController(rice *services.RiceService) *RiceController {
	return &RiceController{
		rice: rice,
	}
}

func (u *RiceController) AddRice(ctx *gin.Context) {
	var req entities.CreateRiceRequest
	if !BindAndValidate(ctx, &req) {
		return
	}
	userID, ok := GetUserID(ctx)
	if !ok {
		return
	}
	count, err := u.rice.AddRice(ctx, userID, &req)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	if count != nil {
		RespondConflict(ctx, http.StatusOK, "rice name already exists")
		return
	}
	RespondSuccess(ctx, nil)
}

func (u *RiceController) GetRiceByUserID(ctx *gin.Context) {
	userID, ok := GetUserID(ctx)
	if !ok {
		return
	}
	resp, err := u.rice.GetListRiceByUserID(ctx, userID)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, resp)
}
