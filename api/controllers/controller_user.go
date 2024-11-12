package controllers

import (
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	user *services.UserService
}

func NewControllerUser(user *services.UserService) *UserController {
	return &UserController{
		user: user,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user entities.CreateUsersRequest
	if !BindAndValidate(ctx, &user) {
		return
	}

	err := u.user.AddUser(ctx, &user)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, nil)
}

func (u *UserController) GetUser(ctx *gin.Context) {
	phoneNumber, ok := GetPhoneNumber(ctx)
	if !ok {
		return
	}
	resp, err := u.user.ProFileUser(ctx, phoneNumber)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, resp)
}
