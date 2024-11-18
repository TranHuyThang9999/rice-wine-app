package controllers

import (
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"

	"github.com/gin-gonic/gin"
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
	userID := GetUserID(ctx)

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
	userID := GetUserID(ctx)

	resp, err := u.rice.GetListRiceByUserID(ctx, userID)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, resp)
}
