package services

import (
	"context"
	"fmt"
	"rice-wine-shop/common/log"
	"rice-wine-shop/core/adapters/interfaces"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/configs"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTService struct {
	user  domain.RepositoryUser
	order *interfaces.OrderClientService
}

func NewJWTService(user domain.RepositoryUser, order *interfaces.OrderClientService) *JWTService {
	return &JWTService{
		user:  user,
		order: order,
	}
}
func (u *JWTService) genToken(ctx context.Context, phoneNumber string, updateAt int64, expireAccess string, role int, userID int64) (string, error) {
	expirationDuration, err := time.ParseDuration(configs.Get().ExpireRefresh)
	if err != nil {
		return "", fmt.Errorf("invalid token expiration duration: %v", err)
	}

	claims := entities.UserClaims{
		UserID:      userID,
		PhoneNumber: phoneNumber,
		UpdateAt:    updateAt,
		Role:        role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.Get().AccessSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func (u *JWTService) Verify(ctx context.Context, tokenString string) (*entities.UserClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.ErrUnexpectedSigningMethod
		}
		return []byte(configs.Get().AccessSecret), nil
	})

	if err != nil || !token.Valid {
		log.Error(err, "error")
		return nil, apperrors.ErrInvalidToken
	}

	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			return nil, apperrors.ErrTokenExpired
		}
	} else {
		return nil, apperrors.ErrInvalidExpiration
	}

	userClaims := &entities.UserClaims{
		PhoneNumber: claims["phoneNumber"].(string),
		UpdateAt:    int64(claims["updatedAt"].(float64)),
	}

	return userClaims, nil
}

func (u *JWTService) Login(ctx context.Context, req *entities.LoginRequest) (*entities.LoginResponse, error) {
	user, err := u.user.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, apperrors.ErrUserNotFound
	}
	token, err := u.genToken(ctx, req.PhoneNumber, user.UpdatedAt, configs.Get().ExpireAccess, user.Role, user.ID)
	if err != nil {
		return nil, err
	}
	referToken, err := u.genToken(ctx, req.PhoneNumber, user.UpdatedAt, configs.Get().ExpireAccess, user.Role, user.ID)
	if err != nil {
		return nil, err
	}

	return &entities.LoginResponse{
		Role:       user.Role,
		Token:      token,
		ReferToken: referToken,
		CreatedAt:  time.Now(),
	}, nil
}

func (u *JWTService) Logout(ctx context.Context, tokenString string) error {
	return nil
}
