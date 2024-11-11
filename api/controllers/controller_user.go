package controllers

import (
	"net/http"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"

	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	user *services.ServiceUser
}

func NewControllerUser(user *services.ServiceUser) *ControllerUser {
	return &ControllerUser{
		user: user,
	}
}

func (u *ControllerUser) CreateUser(ctx *gin.Context) {
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
