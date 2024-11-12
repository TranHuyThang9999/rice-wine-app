package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Token      string    `json:"token"`
	ReferToken string    `json:"referToken"`
	CreatedAt  time.Time `json:"createdAt"`
}

type UserClaims struct {
	PhoneNumber string `json:"phoneNumber"`
	UpdateAt    int64  `json:"updatedAt"`
	Role        int    `json:"role"`
	jwt.RegisteredClaims
}
