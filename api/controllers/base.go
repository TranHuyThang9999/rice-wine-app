package controllers

import (
	"net/http"
	"rice-wine-shop/common/log"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondError(ctx *gin.Context, statusCode int, err error) {
	response := ErrorResponse{
		Error:   true,
		Message: err.Error(),
	}
	ctx.JSON(statusCode, response)
}

func RespondSuccess(ctx *gin.Context, data interface{}) {
	response := SuccessResponse{
		Success: true,
		Message: "success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func BindAndValidate(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		log.Error(err, "error request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}
