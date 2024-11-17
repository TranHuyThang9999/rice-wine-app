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

func RespondConflict(ctx *gin.Context, statusCode int, message string) {
	response := ErrorResponse{
		Error:   true,
		Message: message,
	}
	ctx.JSON(statusCode, response)
}

func BindAndValidate(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		log.Error(err, "error request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

func GetUserID(c *gin.Context) (int64, bool) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(401, gin.H{
			"error": "User ID not found in context",
		})
		return 0, false
	}

	id, ok := userID.(int64)
	if !ok {
		c.JSON(400, gin.H{
			"error": "User ID has an invalid data type",
		})
		return 0, false
	}
	return id, true
}

func GetPhoneNumber(ctx *gin.Context) (string, bool) {
	phoneNumber, ok := ctx.Get("phoneNumber")
	if !ok {
		ctx.JSON(401, gin.H{
			"error": "Phone number not found in context",
		})
		return "", false
	}

	phone, ok := phoneNumber.(string)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "Phone number has an invalid data type",
		})
		return "", false
	}

	return phone, true
}
func GetRole(ctx *gin.Context) (int, bool) {
	role, ok := ctx.Get("role")
	if !ok {
		ctx.JSON(401, gin.H{
			"error": "Phone number not found in context",
		})
		return 0, false
	}

	roleNumber, ok := role.(int)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "Phone number has an invalid data type",
		})
		return 0, false
	}

	return roleNumber, true
}
