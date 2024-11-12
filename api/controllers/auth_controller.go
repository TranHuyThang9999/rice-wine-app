package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"
)

type AuthController struct {
	auth *services.JWTService
}

func NewAuthController(auth *services.JWTService) *AuthController {
	return &AuthController{
		auth: auth,
	}
}

func (u *AuthController) Login(ctx *gin.Context) {
	var req entities.LoginRequest
	if !BindAndValidate(ctx, &req) {
		return
	}
	resp, err := u.auth.Login(ctx, &req)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
	}
	RespondSuccess(ctx, resp)
}
